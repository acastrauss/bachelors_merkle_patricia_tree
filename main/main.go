package main

import "bachelors.com/models"

func main() {

	mpt := models.CreateMPT(models.NodeKey{Key: "a711355"}, models.NodeValue{Value: 1})
	mpt.InsertKVPair(models.NodeKey{Key: "a7ad337"}, models.NodeValue{Value: 22})
	mpt.InsertKVPair(models.NodeKey{Key: "a4ad337"}, models.NodeValue{Value: 322})
	mpt.InsertKVPair(models.NodeKey{Key: "a7ad567"}, models.NodeValue{Value: 32})
	mpt.InsertKVPair(models.NodeKey{Key: "a4a4337"}, models.NodeValue{Value: 33})
}
