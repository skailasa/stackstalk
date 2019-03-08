package main

var StackExchangeClient Client

func init() {
	StackExchangeClient.Host="api.stackexchange.com"
	StackExchangeClient.BasePath="2.2"
	StackExchangeClient.Protocol="https"
}

type Client struct {
	Host, BasePath, Protocol string
	Model Model
}

type requestBody map[string]string

func (p *Client) constructRequest() requestBody {
	r := make(requestBody)
	r["body"] = p.Model.Query
	r["verb"] = p.Model.Verb
	r["stack"] = p.Model.Stack
	return r
}

func (p *Client) GetRequest() {
	r := p.constructRequest()
	println(r["body"])
}
