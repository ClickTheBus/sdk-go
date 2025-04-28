package request

func GetJobStatus(jobId string) (int, []byte, error) {
	return Send("GET", "/jobs/"+jobId+"/", nil)
}
