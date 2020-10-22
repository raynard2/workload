package controller

import (
	"Proto/github.com/binary"
	grpc_from0 "Proto/github.com/monkrus/grpc-from0"
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func P(c echo.Context) error {

	data, workload := DataSet(c)
	//size := strconv.Itoa(int(workload.Batch_Size))

	var Bench int8
	switch workload.WorkloadMetric {
	case "CPU":
		Bench = 0
		//serverRFW.Benchmark_Type = "CPUUtilization_Average"
	case "NETIN":
		Bench = 1
		//serverRFW.Benchmark_Type = "NetworkIn_Average"
	case "NETOUT":
		Bench = 2
		//serverRFW.Benchmark_Type = "NetworkOut_Average"
	case "MEMUTI":
		Bench = 3
		//serverRFW.Benchmark_Type = "MemoryUtilization_Average"
	default:
		Bench = 4
		//serverRFW.Benchmark_Type = "Final_Target"
	}

	fmt.Println(Bench)

	batchSize := (workload.BatchSize / workload.BatchUnit) - 1

	var batch3 []*grpc_from0.Batch
	for i := batchSize; i > (batchSize)-workload.BatchID; i-- {
		//fmt.Println(i)
		f := batchSize - i

		var sam []*grpc_from0.Sample

		c := i * workload.BatchUnit

		for j := c; j < c+workload.BatchUnit; j++ {
			samp := &grpc_from0.Sample{
				CPUUtilization: data[j].CpuUtilizationAverage,
			}
			sam = append(sam, samp)
		}

		batch := &grpc_from0.Batch{
			Batch_ID: int32(f + 1),
			Samples:  sam,
		}
		batch3 = append(batch3, batch)
	}
	//c.JSONPretty(200, batch3, "!")
	//fmt.Println(batch3)
	return Proto(c, workload.RFWID, int32(workload.BatchID), batch3)

}

func Proto(c echo.Context, RFWID string, LASTBATCHID int32, batch []*grpc_from0.Batch) error {

	data, err := binary.EncodeProto(RFWID, LASTBATCHID, batch)

	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	file := &grpc_from0.RFD{}
	binary.DecodeProto(data, file)
	fmt.Print(file)

		c.String(http.StatusOK, string(data))
	return c.JSONBlob(200, data)
}
