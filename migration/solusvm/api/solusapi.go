package api

import (

  "errors"
  "strings"

  "github.com/megamsys/solusvm_go/solusvm"
  "github.com/megamsys/megdcui/automation"
)
const (
  Success = "success"
  InvalidIp = "Invalid ipaddress"
  InvalidIdKey = "Invalid id or key"
)
type SolusClient struct {
  client *solusvm.Client
}

func (s *SolusClient) GetNodes(h *automation.HostInfo) error {
  s.client = solusvm.NewClient(nil,"https://"+ h.SolusMaster + ":5656/api/admin/command.php")
  nodes, _, err := s.client.SolusNodes.ListNodes(map[string]string{ "id":""+ h.Id +"","key":""+ h.Key +"","rdtype": "json", "type":"kvm"})
  if err != nil {
    return err
  }
  h.NodeIds = strings.Split(*nodes.Nodes,",")
  for i := range h.NodeIds {
    s.client = solusvm.NewClient(nil,"https://"+ h.SolusMaster + ":5656/api/admin/command.php")
    nodeinfo, _, Err := s.client.SolusNodes.NodeInfo(map[string]string{ "id":""+ h.Id +"","key":""+ h.Key +"","rdtype": "json", "nodeid":""+ h.NodeIds[i] +""})
    if Err != nil {
      return Err
    }

    if CheckStatus(*nodeinfo.Status, *nodeinfo.Statusmsg) == nil {
      if *nodeinfo.NodeIp == h.SolusNode {
        h.NodeId = h.NodeIds[i]
        return nil
      }
    }
  }

  return errors.New("Can not Find given node ip")
}

func (s *SolusClient) GetVirtualMachines(h *automation.HostInfo) error {
  nodeid := h.NodeId
  s.client = solusvm.NewClient(nil,"https://"+ h.SolusMaster + ":5656/api/admin/command.php")
  servers, _, err := s.client.VirtualServers.ListAllVMs(map[string]string{ "id":""+ h.Id +"","key":""+ h.Key +"","rdtype": "json", "nodeid": ""+ nodeid +""})
  if err != nil {
    return err
  }
  return CheckStatus(*servers.Status, *servers.Statusmsg)
}

func (s *SolusClient) GetClients(h *automation.HostInfo) error {
  s.client = solusvm.NewClient(nil,"https://"+ h.SolusMaster + ":5656/api/admin/command.php")
  clients, _, err := s.client.SolusClients.ListAllClients(map[string]string{ "id":""+ h.Id +"","key":""+ h.Key +"","rdtype": "json"})
  if err != nil {
    return err
  }

  if CheckStatus(*clients.Status, *clients.Statusmsg) != nil {
    return err
  }
  err = storeAccounts(clients)
  if err != nil {
    return err
  }
  return nil
}

func CheckStatus(status, statusmsg string) error {
  if status != Success {
    if statusmsg == InvalidIp {
      return errors.New(InvalidIp)
    }
    if statusmsg == InvalidIdKey {
      return errors.New(InvalidIdKey)
    }
  }
  return nil
}
