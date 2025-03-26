package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RequestBody struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables"` // Use `map[string]interface{}` if it's a dynamic object
}

func main() {
	query := `
		mutation createProductAsynchronous($productSet: ProductSetInput!, $synchronous: Boolean!) {
  productSet(synchronous: $synchronous, input: $productSet) {
    product {
      id
    }
    productSetOperation {
      id
      status
      userErrors {
        code
        field
        message
      }
    }
    userErrors {
      code
      field
      message
    }
  }
}
	`
	variables := `
	{
    "synchronous": false,
    "productSet": {
        "title": "Winter hat",
        "productOptions": [
            {
                "name": "Color",
                "position": 1,
                "values": [
                    {
                        "name": "Grey"
                    },
                    {
                        "name": "Black"
                    }
                ]
            }
        ],
        "variants": [
            {
                "optionValues": [
                    {
                        "optionName": "Color",
                        "name": "Grey"
                    }
                ],
                "price": 79.99
            },
            {
                "optionValues": [
                    {
                        "optionName": "Color",
                        "name": "Black"
                    }
                ],
                "price": 69.99
            }
        ]
    }
}
	`

	if json.Valid([]byte(variables)) {
		fmt.Println("Valid JSON")
	} else {
		fmt.Println("Invalid JSON")
	}
	requestUrl := "https://quick-start-00114f6e.myshopify.com/admin/api/2025-04/graphql.json"
	xShopAt := ""
	jBody := RequestBody{
		Query:     query,
		Variables: json.RawMessage(variables),
	}
	jsonBody, err := json.Marshal(jBody)
	if err != nil {
		fmt.Println("Error marshaling request body:", err)
		return
	}
	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}
	req.Header.Set("X-Shopify-Access-Token", xShopAt)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	// Print response status
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response", string(body))

}
