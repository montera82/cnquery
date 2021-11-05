package resources

import (
	"strconv"

	"github.com/miekg/dns"
	"go.mondoo.io/mondoo/lumi/resources/dnsshake"
)

func (d *lumiDns) id() (string, error) {
	id, _ := d.Fqdn()
	return "dns/" + id, nil
}

func (d *lumiDns) GetParams() (interface{}, error) {
	fqdn, err := d.Fqdn()
	if err != nil {
		return nil, err
	}

	dnsShaker, err := dnsshake.New(fqdn)
	if err != nil {
		return nil, err
	}

	records, err := dnsShaker.Query()
	if err != nil {
		return nil, err
	}

	return jsonToDict(records)
}

// GetRecords returns successful dns records
func (d *lumiDns) GetRecords(params map[string]interface{}) ([]interface{}, error) {
	// convert responses to dns types
	dnsEntries := []interface{}{}
	for k := range params {
		r := params[k].(map[string]interface{})

		// filter by successful dns records
		if r["rCode"] != dns.RcodeToString[dns.RcodeSuccess] {
			continue
		}

		lumiDnsRecord, err := d.Runtime.CreateResource("dns.record",
			"name", r["name"],
			"ttl", r["TTL"],
			"class", r["class"],
			"type", r["type"],
			"rdata", r["rData"],
		)
		if err != nil {
			return nil, err
		}

		dnsEntries = append(dnsEntries, lumiDnsRecord.(DnsRecord))
	}

	return dnsEntries, nil
}

func (d *lumiDnsRecord) id() (string, error) {
	name, _ := d.Name()
	t, _ := d.Type()
	c, _ := d.Class()
	return "dns.record/" + name + "/" + c + "/" + t, nil
}

func (d *lumiDns) GetMx(params map[string]interface{}) ([]interface{}, error) {
	mxEntries := []interface{}{}
	record, ok := params["MX"]
	if !ok {
		return mxEntries, nil
	}

	r := record.(map[string]interface{})

	var name, c, t string
	var ttl int64
	var rdata []interface{}

	if r["name"] != nil {
		name = r["name"].(string)
	}

	if r["class"] != nil {
		c = r["class"].(string)
	}

	if r["type"] != nil {
		t = r["type"].(string)
	}

	if r["TTL"] != nil {
		ttl = r["TTL"].(int64)
	}

	if r["rData"] != nil {
		rdata = r["rData"].([]interface{})
	}

	for j := range rdata {
		entry := rdata[j].(string)

		// use dns package to parse mx entry
		s := name + "\t" + strconv.FormatInt(ttl, 10) + "\t" + c + "\t" + t + "\t" + entry
		got, err := dns.NewRR(s)
		if err != nil {
			return nil, err
		}

		switch v := got.(type) {
		case *dns.MX:
			mxEntry, err := d.Runtime.CreateResource("dns.mxRecord",
				"name", name,
				"preference", int64(v.Preference),
				"domainName", v.Mx,
			)
			if err != nil {
				return nil, err
			}
			mxEntries = append(mxEntries, mxEntry)
		}
	}

	return mxEntries, nil
}

func (d *lumiDnsMxRecord) id() (string, error) {
	name, err := d.Name()
	domainName, _ := d.DomainName()
	return "dns.mx/" + name + "+" + domainName, err
}