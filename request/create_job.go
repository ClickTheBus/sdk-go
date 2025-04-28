package request

func CreateJob(jobType string, payload []byte) (int, []byte, error) {
	return Send("POST", "/jobs/"+jobType+"/", &payload)
}
