package producer

//NewResponseProducer ....
func NewResponseProducer() *ResponseProducer {
	newResponse := &ResponseProducer{}

	newResponse.Topic = "__euka_response_error"
	newResponse.Partition = -999
	newResponse.Offset = -999

	return newResponse
}
