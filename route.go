package restvirt

import (
	"bytes"
	"encoding/json"
	"strings"
)

type Route struct {
	Destination string   `json:"destination"`
	Namespace   string   `json:"namespace"`
	Gateways    []string `json:"gateways"`
}

func (r Route) ID() string {
	return strings.ReplaceAll(r.Destination, "/", "-")
}

var routePath = "routes"

func (c *Client) GetRoutes() ([]*Route, error) {
	resp, err := c.doRequest("GET", nil, routePath)
	if err != nil {
		return nil, err
	}

	err = checkForErrors(resp)
	if err != nil {
		return nil, err
	}

	var routes map[string][]*Route
	err = json.NewDecoder(resp.Body).Decode(&routes)
	if err != nil {
		return nil, err
	}

	return routes["routes"], nil
}

func (c *Client) GetRoutesInNamespace(namespace string) ([]*Route, error) {
	resp, err := c.doRequestWithQueryParams("GET", nil, []KeyValue{
		{Key: "namespace", Value: namespace},
	}, routePath)
	if err != nil {
		return nil, err
	}

	err = checkForErrors(resp)
	if err != nil {
		return nil, err
	}

	var routes map[string][]*Route
	err = json.NewDecoder(resp.Body).Decode(&routes)
	if err != nil {
		return nil, err
	}

	return routes["routes"], nil
}

func (c *Client) GetRoute(destination string) (*Route, error) {
	resp, err := c.doRequest("GET", nil, routePath, destination)
	if err != nil {
		return nil, err
	}

	err = checkForErrors(resp)
	if err != nil {
		return nil, err
	}

	var route Route
	err = json.NewDecoder(resp.Body).Decode(&route)
	if err != nil {
		return nil, err
	}

	return &route, nil
}

func (c *Client) SetRoute(route Route) error {
	routeJSON, err := json.Marshal(route)
	if err != nil {
		return err
	}

	resp, err := c.doRequest("PUT", bytes.NewBuffer(routeJSON), routePath, route.ID())
	if err != nil {
		return err
	}

	err = checkForErrors(resp)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteRoute(destination string) error {
	resp, err := c.doRequest("DELETE", nil, routePath, destination)
	if err != nil {
		return err
	}

	err = checkForErrors(resp)
	if err != nil {
		return err
	}

	return nil
}
