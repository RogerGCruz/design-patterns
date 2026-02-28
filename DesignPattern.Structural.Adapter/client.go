package main

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client: Insert Lightning Connector into Computer")
	com.InsertIntoLightningPort()
}
