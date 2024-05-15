package pkg

import "fmt"

func CreateResources(siteName string, deploymentName string, serviceName string, ingressName string, hostName string) {
	//deployment作成
	CreateDeployment(siteName, deploymentName)
	//service作成
	CreateService(siteName, serviceName)
	//ingress作成
	CreateIngress(ingressName, hostName, serviceName)
}

func AddDnsResouces(RecordName, RecordType, Ttl, Content string) error {
	//レコード追加
	resp, err := AddRecords(RecordName, RecordType, Ttl, Content)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
	return nil
}

func DeleteResources(deploymentName string, serviceName string, ingressName string) {
	//deployment削除
	DeleteDeployment(deploymentName)
	//service削除
	DeleteService(serviceName)
	//ingress削除
	DeleteIngress(ingressName)
}

func DeleteDnsResouces(deleteRecordName string) error {
	resp, err := DeleteRecords(deleteRecordName)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
	return nil
}
