package controllers

import (
	"encoding/base64"
	"log"

	"github.com/gin-gonic/gin"
)

//Root endpoint
func GetRoot(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Base64-Api Created With ❤️ by Ujjwal Mahar",
	})}

//For encoding the value 
func PostEncode(c *gin.Context) {

	//Getting input by the user
	var Encoder struct{
		EncodeValue string  
	}
	c.Bind(&Encoder)

	//Encoder logic
	encoder:= base64.StdEncoding.EncodeToString([]byte(Encoder.EncodeValue))
	
	//return response

	c.JSON(200, gin.H{
		"EncodedResult": encoder,
	})
}

//for decoding the Base64 vale 

func PostDecode(c *gin.Context){

	//get data from the user 
	var Decoder struct{
		ToDecodeValue string
	}
	c.Bind(&Decoder)

	//decode logic

	decoder,err := base64.StdEncoding.DecodeString(Decoder.ToDecodeValue)
	if err !=nil{
		log.Fatal("Error decoding the value")
	}

	//return the response
	c.JSON(200, gin.H{
		"DecodedResult": string(decoder),
	})

}

