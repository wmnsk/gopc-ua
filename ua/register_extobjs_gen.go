// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Code generated by cmd/service. DO NOT EDIT!

package ua

import (
	"github.com/liuxgo/opcua/id"
)

func init() {
	RegisterExtensionObject(NewNumericNodeID(0, id.KeyValuePair_Encoding_DefaultBinary), new(KeyValuePair))
	RegisterExtensionObject(NewNumericNodeID(0, id.AdditionalParametersType_Encoding_DefaultBinary), new(AdditionalParametersType))
	RegisterExtensionObject(NewNumericNodeID(0, id.EphemeralKeyType_Encoding_DefaultBinary), new(EphemeralKeyType))
	RegisterExtensionObject(NewNumericNodeID(0, id.EndpointType_Encoding_DefaultBinary), new(EndpointType))
	RegisterExtensionObject(NewNumericNodeID(0, id.IdentityMappingRuleType_Encoding_DefaultBinary), new(IdentityMappingRuleType))
	RegisterExtensionObject(NewNumericNodeID(0, id.TrustListDataType_Encoding_DefaultBinary), new(TrustListDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.DecimalDataType_Encoding_DefaultBinary), new(DecimalDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.DataTypeSchemaHeader_Encoding_DefaultBinary), new(DataTypeSchemaHeader))
	RegisterExtensionObject(NewNumericNodeID(0, id.DataTypeDescription_Encoding_DefaultBinary), new(DataTypeDescription))
	RegisterExtensionObject(NewNumericNodeID(0, id.StructureDescription_Encoding_DefaultBinary), new(StructureDescription))
	RegisterExtensionObject(NewNumericNodeID(0, id.EnumDescription_Encoding_DefaultBinary), new(EnumDescription))
	RegisterExtensionObject(NewNumericNodeID(0, id.SimpleTypeDescription_Encoding_DefaultBinary), new(SimpleTypeDescription))
	RegisterExtensionObject(NewNumericNodeID(0, id.UABinaryFileDataType_Encoding_DefaultBinary), new(UABinaryFileDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.DataSetMetaDataType_Encoding_DefaultBinary), new(DataSetMetaDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.FieldMetaData_Encoding_DefaultBinary), new(FieldMetaData))
	RegisterExtensionObject(NewNumericNodeID(0, id.ConfigurationVersionDataType_Encoding_DefaultBinary), new(ConfigurationVersionDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.PublishedDataSetDataType_Encoding_DefaultBinary), new(PublishedDataSetDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.PublishedDataSetSourceDataType_Encoding_DefaultBinary), new(PublishedDataSetSourceDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.PublishedVariableDataType_Encoding_DefaultBinary), new(PublishedVariableDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.PublishedDataItemsDataType_Encoding_DefaultBinary), new(PublishedDataItemsDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.PublishedEventsDataType_Encoding_DefaultBinary), new(PublishedEventsDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.DataSetWriterDataType_Encoding_DefaultBinary), new(DataSetWriterDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.DataSetWriterTransportDataType_Encoding_DefaultBinary), new(DataSetWriterTransportDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.DataSetWriterMessageDataType_Encoding_DefaultBinary), new(DataSetWriterMessageDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.PubSubGroupDataType_Encoding_DefaultBinary), new(PubSubGroupDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.WriterGroupDataType_Encoding_DefaultBinary), new(WriterGroupDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.WriterGroupTransportDataType_Encoding_DefaultBinary), new(WriterGroupTransportDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.WriterGroupMessageDataType_Encoding_DefaultBinary), new(WriterGroupMessageDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.PubSubConnectionDataType_Encoding_DefaultBinary), new(PubSubConnectionDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.ConnectionTransportDataType_Encoding_DefaultBinary), new(ConnectionTransportDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.NetworkAddressDataType_Encoding_DefaultBinary), new(NetworkAddressDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.NetworkAddressURLDataType_Encoding_DefaultBinary), new(NetworkAddressURLDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.ReaderGroupDataType_Encoding_DefaultBinary), new(ReaderGroupDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.ReaderGroupTransportDataType_Encoding_DefaultBinary), new(ReaderGroupTransportDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.ReaderGroupMessageDataType_Encoding_DefaultBinary), new(ReaderGroupMessageDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.DataSetReaderDataType_Encoding_DefaultBinary), new(DataSetReaderDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.DataSetReaderTransportDataType_Encoding_DefaultBinary), new(DataSetReaderTransportDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.DataSetReaderMessageDataType_Encoding_DefaultBinary), new(DataSetReaderMessageDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.SubscribedDataSetDataType_Encoding_DefaultBinary), new(SubscribedDataSetDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.TargetVariablesDataType_Encoding_DefaultBinary), new(TargetVariablesDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.FieldTargetDataType_Encoding_DefaultBinary), new(FieldTargetDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.SubscribedDataSetMirrorDataType_Encoding_DefaultBinary), new(SubscribedDataSetMirrorDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.PubSubConfigurationDataType_Encoding_DefaultBinary), new(PubSubConfigurationDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.UADPWriterGroupMessageDataType_Encoding_DefaultBinary), new(UADPWriterGroupMessageDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.UADPDataSetWriterMessageDataType_Encoding_DefaultBinary), new(UADPDataSetWriterMessageDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.UADPDataSetReaderMessageDataType_Encoding_DefaultBinary), new(UADPDataSetReaderMessageDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.JSONWriterGroupMessageDataType_Encoding_DefaultBinary), new(JSONWriterGroupMessageDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.JSONDataSetWriterMessageDataType_Encoding_DefaultBinary), new(JSONDataSetWriterMessageDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.JSONDataSetReaderMessageDataType_Encoding_DefaultBinary), new(JSONDataSetReaderMessageDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.DatagramConnectionTransportDataType_Encoding_DefaultBinary), new(DatagramConnectionTransportDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.DatagramWriterGroupTransportDataType_Encoding_DefaultBinary), new(DatagramWriterGroupTransportDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.BrokerConnectionTransportDataType_Encoding_DefaultBinary), new(BrokerConnectionTransportDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.BrokerWriterGroupTransportDataType_Encoding_DefaultBinary), new(BrokerWriterGroupTransportDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.BrokerDataSetWriterTransportDataType_Encoding_DefaultBinary), new(BrokerDataSetWriterTransportDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.BrokerDataSetReaderTransportDataType_Encoding_DefaultBinary), new(BrokerDataSetReaderTransportDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.RolePermissionType_Encoding_DefaultBinary), new(RolePermissionType))
	RegisterExtensionObject(NewNumericNodeID(0, id.StructureField_Encoding_DefaultBinary), new(StructureField))
	RegisterExtensionObject(NewNumericNodeID(0, id.StructureDefinition_Encoding_DefaultBinary), new(StructureDefinition))
	RegisterExtensionObject(NewNumericNodeID(0, id.EnumDefinition_Encoding_DefaultBinary), new(EnumDefinition))
	RegisterExtensionObject(NewNumericNodeID(0, id.Node_Encoding_DefaultBinary), new(Node))
	RegisterExtensionObject(NewNumericNodeID(0, id.InstanceNode_Encoding_DefaultBinary), new(InstanceNode))
	RegisterExtensionObject(NewNumericNodeID(0, id.TypeNode_Encoding_DefaultBinary), new(TypeNode))
	RegisterExtensionObject(NewNumericNodeID(0, id.ObjectNode_Encoding_DefaultBinary), new(ObjectNode))
	RegisterExtensionObject(NewNumericNodeID(0, id.ObjectTypeNode_Encoding_DefaultBinary), new(ObjectTypeNode))
	RegisterExtensionObject(NewNumericNodeID(0, id.VariableNode_Encoding_DefaultBinary), new(VariableNode))
	RegisterExtensionObject(NewNumericNodeID(0, id.VariableTypeNode_Encoding_DefaultBinary), new(VariableTypeNode))
	RegisterExtensionObject(NewNumericNodeID(0, id.ReferenceTypeNode_Encoding_DefaultBinary), new(ReferenceTypeNode))
	RegisterExtensionObject(NewNumericNodeID(0, id.MethodNode_Encoding_DefaultBinary), new(MethodNode))
	RegisterExtensionObject(NewNumericNodeID(0, id.ViewNode_Encoding_DefaultBinary), new(ViewNode))
	RegisterExtensionObject(NewNumericNodeID(0, id.DataTypeNode_Encoding_DefaultBinary), new(DataTypeNode))
	RegisterExtensionObject(NewNumericNodeID(0, id.ReferenceNode_Encoding_DefaultBinary), new(ReferenceNode))
	RegisterExtensionObject(NewNumericNodeID(0, id.Argument_Encoding_DefaultBinary), new(Argument))
	RegisterExtensionObject(NewNumericNodeID(0, id.EnumValueType_Encoding_DefaultBinary), new(EnumValueType))
	RegisterExtensionObject(NewNumericNodeID(0, id.EnumField_Encoding_DefaultBinary), new(EnumField))
	RegisterExtensionObject(NewNumericNodeID(0, id.OptionSet_Encoding_DefaultBinary), new(OptionSet))
	RegisterExtensionObject(NewNumericNodeID(0, id.Union_Encoding_DefaultBinary), new(Union))
	RegisterExtensionObject(NewNumericNodeID(0, id.TimeZoneDataType_Encoding_DefaultBinary), new(TimeZoneDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.ApplicationDescription_Encoding_DefaultBinary), new(ApplicationDescription))
	RegisterExtensionObject(NewNumericNodeID(0, id.RequestHeader_Encoding_DefaultBinary), new(RequestHeader))
	RegisterExtensionObject(NewNumericNodeID(0, id.ResponseHeader_Encoding_DefaultBinary), new(ResponseHeader))
	RegisterExtensionObject(NewNumericNodeID(0, id.ServiceFault_Encoding_DefaultBinary), new(ServiceFault))
	RegisterExtensionObject(NewNumericNodeID(0, id.SessionlessInvokeRequestType_Encoding_DefaultBinary), new(SessionlessInvokeRequestType))
	RegisterExtensionObject(NewNumericNodeID(0, id.SessionlessInvokeResponseType_Encoding_DefaultBinary), new(SessionlessInvokeResponseType))
	RegisterExtensionObject(NewNumericNodeID(0, id.FindServersRequest_Encoding_DefaultBinary), new(FindServersRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.FindServersResponse_Encoding_DefaultBinary), new(FindServersResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.ServerOnNetwork_Encoding_DefaultBinary), new(ServerOnNetwork))
	RegisterExtensionObject(NewNumericNodeID(0, id.FindServersOnNetworkRequest_Encoding_DefaultBinary), new(FindServersOnNetworkRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.FindServersOnNetworkResponse_Encoding_DefaultBinary), new(FindServersOnNetworkResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.UserTokenPolicy_Encoding_DefaultBinary), new(UserTokenPolicy))
	RegisterExtensionObject(NewNumericNodeID(0, id.EndpointDescription_Encoding_DefaultBinary), new(EndpointDescription))
	RegisterExtensionObject(NewNumericNodeID(0, id.GetEndpointsRequest_Encoding_DefaultBinary), new(GetEndpointsRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.GetEndpointsResponse_Encoding_DefaultBinary), new(GetEndpointsResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.RegisteredServer_Encoding_DefaultBinary), new(RegisteredServer))
	RegisterExtensionObject(NewNumericNodeID(0, id.RegisterServerRequest_Encoding_DefaultBinary), new(RegisterServerRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.RegisterServerResponse_Encoding_DefaultBinary), new(RegisterServerResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.DiscoveryConfiguration_Encoding_DefaultBinary), new(DiscoveryConfiguration))
	RegisterExtensionObject(NewNumericNodeID(0, id.MdnsDiscoveryConfiguration_Encoding_DefaultBinary), new(MdnsDiscoveryConfiguration))
	RegisterExtensionObject(NewNumericNodeID(0, id.RegisterServer2Request_Encoding_DefaultBinary), new(RegisterServer2Request))
	RegisterExtensionObject(NewNumericNodeID(0, id.RegisterServer2Response_Encoding_DefaultBinary), new(RegisterServer2Response))
	RegisterExtensionObject(NewNumericNodeID(0, id.ChannelSecurityToken_Encoding_DefaultBinary), new(ChannelSecurityToken))
	RegisterExtensionObject(NewNumericNodeID(0, id.OpenSecureChannelRequest_Encoding_DefaultBinary), new(OpenSecureChannelRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.OpenSecureChannelResponse_Encoding_DefaultBinary), new(OpenSecureChannelResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.CloseSecureChannelRequest_Encoding_DefaultBinary), new(CloseSecureChannelRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.CloseSecureChannelResponse_Encoding_DefaultBinary), new(CloseSecureChannelResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.SignedSoftwareCertificate_Encoding_DefaultBinary), new(SignedSoftwareCertificate))
	RegisterExtensionObject(NewNumericNodeID(0, id.SignatureData_Encoding_DefaultBinary), new(SignatureData))
	RegisterExtensionObject(NewNumericNodeID(0, id.CreateSessionRequest_Encoding_DefaultBinary), new(CreateSessionRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.CreateSessionResponse_Encoding_DefaultBinary), new(CreateSessionResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.UserIdentityToken_Encoding_DefaultBinary), new(UserIdentityToken))
	RegisterExtensionObject(NewNumericNodeID(0, id.AnonymousIdentityToken_Encoding_DefaultBinary), new(AnonymousIdentityToken))
	RegisterExtensionObject(NewNumericNodeID(0, id.UserNameIdentityToken_Encoding_DefaultBinary), new(UserNameIdentityToken))
	RegisterExtensionObject(NewNumericNodeID(0, id.X509IdentityToken_Encoding_DefaultBinary), new(X509IdentityToken))
	RegisterExtensionObject(NewNumericNodeID(0, id.IssuedIdentityToken_Encoding_DefaultBinary), new(IssuedIdentityToken))
	RegisterExtensionObject(NewNumericNodeID(0, id.ActivateSessionRequest_Encoding_DefaultBinary), new(ActivateSessionRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.ActivateSessionResponse_Encoding_DefaultBinary), new(ActivateSessionResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.CloseSessionRequest_Encoding_DefaultBinary), new(CloseSessionRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.CloseSessionResponse_Encoding_DefaultBinary), new(CloseSessionResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.CancelRequest_Encoding_DefaultBinary), new(CancelRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.CancelResponse_Encoding_DefaultBinary), new(CancelResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.NodeAttributes_Encoding_DefaultBinary), new(NodeAttributes))
	RegisterExtensionObject(NewNumericNodeID(0, id.ObjectAttributes_Encoding_DefaultBinary), new(ObjectAttributes))
	RegisterExtensionObject(NewNumericNodeID(0, id.VariableAttributes_Encoding_DefaultBinary), new(VariableAttributes))
	RegisterExtensionObject(NewNumericNodeID(0, id.MethodAttributes_Encoding_DefaultBinary), new(MethodAttributes))
	RegisterExtensionObject(NewNumericNodeID(0, id.ObjectTypeAttributes_Encoding_DefaultBinary), new(ObjectTypeAttributes))
	RegisterExtensionObject(NewNumericNodeID(0, id.VariableTypeAttributes_Encoding_DefaultBinary), new(VariableTypeAttributes))
	RegisterExtensionObject(NewNumericNodeID(0, id.ReferenceTypeAttributes_Encoding_DefaultBinary), new(ReferenceTypeAttributes))
	RegisterExtensionObject(NewNumericNodeID(0, id.DataTypeAttributes_Encoding_DefaultBinary), new(DataTypeAttributes))
	RegisterExtensionObject(NewNumericNodeID(0, id.ViewAttributes_Encoding_DefaultBinary), new(ViewAttributes))
	RegisterExtensionObject(NewNumericNodeID(0, id.GenericAttributeValue_Encoding_DefaultBinary), new(GenericAttributeValue))
	RegisterExtensionObject(NewNumericNodeID(0, id.GenericAttributes_Encoding_DefaultBinary), new(GenericAttributes))
	RegisterExtensionObject(NewNumericNodeID(0, id.AddNodesItem_Encoding_DefaultBinary), new(AddNodesItem))
	RegisterExtensionObject(NewNumericNodeID(0, id.AddNodesResult_Encoding_DefaultBinary), new(AddNodesResult))
	RegisterExtensionObject(NewNumericNodeID(0, id.AddNodesRequest_Encoding_DefaultBinary), new(AddNodesRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.AddNodesResponse_Encoding_DefaultBinary), new(AddNodesResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.AddReferencesItem_Encoding_DefaultBinary), new(AddReferencesItem))
	RegisterExtensionObject(NewNumericNodeID(0, id.AddReferencesRequest_Encoding_DefaultBinary), new(AddReferencesRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.AddReferencesResponse_Encoding_DefaultBinary), new(AddReferencesResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.DeleteNodesItem_Encoding_DefaultBinary), new(DeleteNodesItem))
	RegisterExtensionObject(NewNumericNodeID(0, id.DeleteNodesRequest_Encoding_DefaultBinary), new(DeleteNodesRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.DeleteNodesResponse_Encoding_DefaultBinary), new(DeleteNodesResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.DeleteReferencesItem_Encoding_DefaultBinary), new(DeleteReferencesItem))
	RegisterExtensionObject(NewNumericNodeID(0, id.DeleteReferencesRequest_Encoding_DefaultBinary), new(DeleteReferencesRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.DeleteReferencesResponse_Encoding_DefaultBinary), new(DeleteReferencesResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.ViewDescription_Encoding_DefaultBinary), new(ViewDescription))
	RegisterExtensionObject(NewNumericNodeID(0, id.BrowseDescription_Encoding_DefaultBinary), new(BrowseDescription))
	RegisterExtensionObject(NewNumericNodeID(0, id.ReferenceDescription_Encoding_DefaultBinary), new(ReferenceDescription))
	RegisterExtensionObject(NewNumericNodeID(0, id.BrowseResult_Encoding_DefaultBinary), new(BrowseResult))
	RegisterExtensionObject(NewNumericNodeID(0, id.BrowseRequest_Encoding_DefaultBinary), new(BrowseRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.BrowseResponse_Encoding_DefaultBinary), new(BrowseResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.BrowseNextRequest_Encoding_DefaultBinary), new(BrowseNextRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.BrowseNextResponse_Encoding_DefaultBinary), new(BrowseNextResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.RelativePathElement_Encoding_DefaultBinary), new(RelativePathElement))
	RegisterExtensionObject(NewNumericNodeID(0, id.RelativePath_Encoding_DefaultBinary), new(RelativePath))
	RegisterExtensionObject(NewNumericNodeID(0, id.BrowsePath_Encoding_DefaultBinary), new(BrowsePath))
	RegisterExtensionObject(NewNumericNodeID(0, id.BrowsePathTarget_Encoding_DefaultBinary), new(BrowsePathTarget))
	RegisterExtensionObject(NewNumericNodeID(0, id.BrowsePathResult_Encoding_DefaultBinary), new(BrowsePathResult))
	RegisterExtensionObject(NewNumericNodeID(0, id.TranslateBrowsePathsToNodeIDsRequest_Encoding_DefaultBinary), new(TranslateBrowsePathsToNodeIDsRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.TranslateBrowsePathsToNodeIDsResponse_Encoding_DefaultBinary), new(TranslateBrowsePathsToNodeIDsResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.RegisterNodesRequest_Encoding_DefaultBinary), new(RegisterNodesRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.RegisterNodesResponse_Encoding_DefaultBinary), new(RegisterNodesResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.UnregisterNodesRequest_Encoding_DefaultBinary), new(UnregisterNodesRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.UnregisterNodesResponse_Encoding_DefaultBinary), new(UnregisterNodesResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.EndpointConfiguration_Encoding_DefaultBinary), new(EndpointConfiguration))
	RegisterExtensionObject(NewNumericNodeID(0, id.QueryDataDescription_Encoding_DefaultBinary), new(QueryDataDescription))
	RegisterExtensionObject(NewNumericNodeID(0, id.NodeTypeDescription_Encoding_DefaultBinary), new(NodeTypeDescription))
	RegisterExtensionObject(NewNumericNodeID(0, id.QueryDataSet_Encoding_DefaultBinary), new(QueryDataSet))
	RegisterExtensionObject(NewNumericNodeID(0, id.NodeReference_Encoding_DefaultBinary), new(NodeReference))
	RegisterExtensionObject(NewNumericNodeID(0, id.ContentFilterElement_Encoding_DefaultBinary), new(ContentFilterElement))
	RegisterExtensionObject(NewNumericNodeID(0, id.ContentFilter_Encoding_DefaultBinary), new(ContentFilter))
	RegisterExtensionObject(NewNumericNodeID(0, id.FilterOperand_Encoding_DefaultBinary), new(FilterOperand))
	RegisterExtensionObject(NewNumericNodeID(0, id.ElementOperand_Encoding_DefaultBinary), new(ElementOperand))
	RegisterExtensionObject(NewNumericNodeID(0, id.LiteralOperand_Encoding_DefaultBinary), new(LiteralOperand))
	RegisterExtensionObject(NewNumericNodeID(0, id.AttributeOperand_Encoding_DefaultBinary), new(AttributeOperand))
	RegisterExtensionObject(NewNumericNodeID(0, id.SimpleAttributeOperand_Encoding_DefaultBinary), new(SimpleAttributeOperand))
	RegisterExtensionObject(NewNumericNodeID(0, id.ContentFilterElementResult_Encoding_DefaultBinary), new(ContentFilterElementResult))
	RegisterExtensionObject(NewNumericNodeID(0, id.ContentFilterResult_Encoding_DefaultBinary), new(ContentFilterResult))
	RegisterExtensionObject(NewNumericNodeID(0, id.ParsingResult_Encoding_DefaultBinary), new(ParsingResult))
	RegisterExtensionObject(NewNumericNodeID(0, id.QueryFirstRequest_Encoding_DefaultBinary), new(QueryFirstRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.QueryFirstResponse_Encoding_DefaultBinary), new(QueryFirstResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.QueryNextRequest_Encoding_DefaultBinary), new(QueryNextRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.QueryNextResponse_Encoding_DefaultBinary), new(QueryNextResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.ReadValueID_Encoding_DefaultBinary), new(ReadValueID))
	RegisterExtensionObject(NewNumericNodeID(0, id.ReadRequest_Encoding_DefaultBinary), new(ReadRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.ReadResponse_Encoding_DefaultBinary), new(ReadResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.HistoryReadValueID_Encoding_DefaultBinary), new(HistoryReadValueID))
	RegisterExtensionObject(NewNumericNodeID(0, id.HistoryReadResult_Encoding_DefaultBinary), new(HistoryReadResult))
	RegisterExtensionObject(NewNumericNodeID(0, id.HistoryReadDetails_Encoding_DefaultBinary), new(HistoryReadDetails))
	RegisterExtensionObject(NewNumericNodeID(0, id.ReadEventDetails_Encoding_DefaultBinary), new(ReadEventDetails))
	RegisterExtensionObject(NewNumericNodeID(0, id.ReadRawModifiedDetails_Encoding_DefaultBinary), new(ReadRawModifiedDetails))
	RegisterExtensionObject(NewNumericNodeID(0, id.ReadProcessedDetails_Encoding_DefaultBinary), new(ReadProcessedDetails))
	RegisterExtensionObject(NewNumericNodeID(0, id.ReadAtTimeDetails_Encoding_DefaultBinary), new(ReadAtTimeDetails))
	RegisterExtensionObject(NewNumericNodeID(0, id.HistoryData_Encoding_DefaultBinary), new(HistoryData))
	RegisterExtensionObject(NewNumericNodeID(0, id.ModificationInfo_Encoding_DefaultBinary), new(ModificationInfo))
	RegisterExtensionObject(NewNumericNodeID(0, id.HistoryModifiedData_Encoding_DefaultBinary), new(HistoryModifiedData))
	RegisterExtensionObject(NewNumericNodeID(0, id.HistoryEvent_Encoding_DefaultBinary), new(HistoryEvent))
	RegisterExtensionObject(NewNumericNodeID(0, id.HistoryReadRequest_Encoding_DefaultBinary), new(HistoryReadRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.HistoryReadResponse_Encoding_DefaultBinary), new(HistoryReadResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.WriteValue_Encoding_DefaultBinary), new(WriteValue))
	RegisterExtensionObject(NewNumericNodeID(0, id.WriteRequest_Encoding_DefaultBinary), new(WriteRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.WriteResponse_Encoding_DefaultBinary), new(WriteResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.HistoryUpdateDetails_Encoding_DefaultBinary), new(HistoryUpdateDetails))
	RegisterExtensionObject(NewNumericNodeID(0, id.UpdateDataDetails_Encoding_DefaultBinary), new(UpdateDataDetails))
	RegisterExtensionObject(NewNumericNodeID(0, id.UpdateStructureDataDetails_Encoding_DefaultBinary), new(UpdateStructureDataDetails))
	RegisterExtensionObject(NewNumericNodeID(0, id.UpdateEventDetails_Encoding_DefaultBinary), new(UpdateEventDetails))
	RegisterExtensionObject(NewNumericNodeID(0, id.DeleteRawModifiedDetails_Encoding_DefaultBinary), new(DeleteRawModifiedDetails))
	RegisterExtensionObject(NewNumericNodeID(0, id.DeleteAtTimeDetails_Encoding_DefaultBinary), new(DeleteAtTimeDetails))
	RegisterExtensionObject(NewNumericNodeID(0, id.DeleteEventDetails_Encoding_DefaultBinary), new(DeleteEventDetails))
	RegisterExtensionObject(NewNumericNodeID(0, id.HistoryUpdateResult_Encoding_DefaultBinary), new(HistoryUpdateResult))
	RegisterExtensionObject(NewNumericNodeID(0, id.HistoryUpdateRequest_Encoding_DefaultBinary), new(HistoryUpdateRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.HistoryUpdateResponse_Encoding_DefaultBinary), new(HistoryUpdateResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.CallMethodRequest_Encoding_DefaultBinary), new(CallMethodRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.CallMethodResult_Encoding_DefaultBinary), new(CallMethodResult))
	RegisterExtensionObject(NewNumericNodeID(0, id.CallRequest_Encoding_DefaultBinary), new(CallRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.CallResponse_Encoding_DefaultBinary), new(CallResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.MonitoringFilter_Encoding_DefaultBinary), new(MonitoringFilter))
	RegisterExtensionObject(NewNumericNodeID(0, id.DataChangeFilter_Encoding_DefaultBinary), new(DataChangeFilter))
	RegisterExtensionObject(NewNumericNodeID(0, id.EventFilter_Encoding_DefaultBinary), new(EventFilter))
	RegisterExtensionObject(NewNumericNodeID(0, id.AggregateConfiguration_Encoding_DefaultBinary), new(AggregateConfiguration))
	RegisterExtensionObject(NewNumericNodeID(0, id.AggregateFilter_Encoding_DefaultBinary), new(AggregateFilter))
	RegisterExtensionObject(NewNumericNodeID(0, id.MonitoringFilterResult_Encoding_DefaultBinary), new(MonitoringFilterResult))
	RegisterExtensionObject(NewNumericNodeID(0, id.EventFilterResult_Encoding_DefaultBinary), new(EventFilterResult))
	RegisterExtensionObject(NewNumericNodeID(0, id.AggregateFilterResult_Encoding_DefaultBinary), new(AggregateFilterResult))
	RegisterExtensionObject(NewNumericNodeID(0, id.MonitoringParameters_Encoding_DefaultBinary), new(MonitoringParameters))
	RegisterExtensionObject(NewNumericNodeID(0, id.MonitoredItemCreateRequest_Encoding_DefaultBinary), new(MonitoredItemCreateRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.MonitoredItemCreateResult_Encoding_DefaultBinary), new(MonitoredItemCreateResult))
	RegisterExtensionObject(NewNumericNodeID(0, id.CreateMonitoredItemsRequest_Encoding_DefaultBinary), new(CreateMonitoredItemsRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.CreateMonitoredItemsResponse_Encoding_DefaultBinary), new(CreateMonitoredItemsResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.MonitoredItemModifyRequest_Encoding_DefaultBinary), new(MonitoredItemModifyRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.MonitoredItemModifyResult_Encoding_DefaultBinary), new(MonitoredItemModifyResult))
	RegisterExtensionObject(NewNumericNodeID(0, id.ModifyMonitoredItemsRequest_Encoding_DefaultBinary), new(ModifyMonitoredItemsRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.ModifyMonitoredItemsResponse_Encoding_DefaultBinary), new(ModifyMonitoredItemsResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.SetMonitoringModeRequest_Encoding_DefaultBinary), new(SetMonitoringModeRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.SetMonitoringModeResponse_Encoding_DefaultBinary), new(SetMonitoringModeResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.SetTriggeringRequest_Encoding_DefaultBinary), new(SetTriggeringRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.SetTriggeringResponse_Encoding_DefaultBinary), new(SetTriggeringResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.DeleteMonitoredItemsRequest_Encoding_DefaultBinary), new(DeleteMonitoredItemsRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.DeleteMonitoredItemsResponse_Encoding_DefaultBinary), new(DeleteMonitoredItemsResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.CreateSubscriptionRequest_Encoding_DefaultBinary), new(CreateSubscriptionRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.CreateSubscriptionResponse_Encoding_DefaultBinary), new(CreateSubscriptionResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.ModifySubscriptionRequest_Encoding_DefaultBinary), new(ModifySubscriptionRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.ModifySubscriptionResponse_Encoding_DefaultBinary), new(ModifySubscriptionResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.SetPublishingModeRequest_Encoding_DefaultBinary), new(SetPublishingModeRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.SetPublishingModeResponse_Encoding_DefaultBinary), new(SetPublishingModeResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.NotificationMessage_Encoding_DefaultBinary), new(NotificationMessage))
	RegisterExtensionObject(NewNumericNodeID(0, id.NotificationData_Encoding_DefaultBinary), new(NotificationData))
	RegisterExtensionObject(NewNumericNodeID(0, id.DataChangeNotification_Encoding_DefaultBinary), new(DataChangeNotification))
	RegisterExtensionObject(NewNumericNodeID(0, id.MonitoredItemNotification_Encoding_DefaultBinary), new(MonitoredItemNotification))
	RegisterExtensionObject(NewNumericNodeID(0, id.EventNotificationList_Encoding_DefaultBinary), new(EventNotificationList))
	RegisterExtensionObject(NewNumericNodeID(0, id.EventFieldList_Encoding_DefaultBinary), new(EventFieldList))
	RegisterExtensionObject(NewNumericNodeID(0, id.HistoryEventFieldList_Encoding_DefaultBinary), new(HistoryEventFieldList))
	RegisterExtensionObject(NewNumericNodeID(0, id.StatusChangeNotification_Encoding_DefaultBinary), new(StatusChangeNotification))
	RegisterExtensionObject(NewNumericNodeID(0, id.SubscriptionAcknowledgement_Encoding_DefaultBinary), new(SubscriptionAcknowledgement))
	RegisterExtensionObject(NewNumericNodeID(0, id.PublishRequest_Encoding_DefaultBinary), new(PublishRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.PublishResponse_Encoding_DefaultBinary), new(PublishResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.RepublishRequest_Encoding_DefaultBinary), new(RepublishRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.RepublishResponse_Encoding_DefaultBinary), new(RepublishResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.TransferResult_Encoding_DefaultBinary), new(TransferResult))
	RegisterExtensionObject(NewNumericNodeID(0, id.TransferSubscriptionsRequest_Encoding_DefaultBinary), new(TransferSubscriptionsRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.TransferSubscriptionsResponse_Encoding_DefaultBinary), new(TransferSubscriptionsResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.DeleteSubscriptionsRequest_Encoding_DefaultBinary), new(DeleteSubscriptionsRequest))
	RegisterExtensionObject(NewNumericNodeID(0, id.DeleteSubscriptionsResponse_Encoding_DefaultBinary), new(DeleteSubscriptionsResponse))
	RegisterExtensionObject(NewNumericNodeID(0, id.BuildInfo_Encoding_DefaultBinary), new(BuildInfo))
	RegisterExtensionObject(NewNumericNodeID(0, id.RedundantServerDataType_Encoding_DefaultBinary), new(RedundantServerDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.EndpointURLListDataType_Encoding_DefaultBinary), new(EndpointURLListDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.NetworkGroupDataType_Encoding_DefaultBinary), new(NetworkGroupDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.SamplingIntervalDiagnosticsDataType_Encoding_DefaultBinary), new(SamplingIntervalDiagnosticsDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.ServerDiagnosticsSummaryDataType_Encoding_DefaultBinary), new(ServerDiagnosticsSummaryDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.ServerStatusDataType_Encoding_DefaultBinary), new(ServerStatusDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.SessionDiagnosticsDataType_Encoding_DefaultBinary), new(SessionDiagnosticsDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.SessionSecurityDiagnosticsDataType_Encoding_DefaultBinary), new(SessionSecurityDiagnosticsDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.ServiceCounterDataType_Encoding_DefaultBinary), new(ServiceCounterDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.StatusResult_Encoding_DefaultBinary), new(StatusResult))
	RegisterExtensionObject(NewNumericNodeID(0, id.SubscriptionDiagnosticsDataType_Encoding_DefaultBinary), new(SubscriptionDiagnosticsDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.ModelChangeStructureDataType_Encoding_DefaultBinary), new(ModelChangeStructureDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.SemanticChangeStructureDataType_Encoding_DefaultBinary), new(SemanticChangeStructureDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.Range_Encoding_DefaultBinary), new(Range))
	RegisterExtensionObject(NewNumericNodeID(0, id.EUInformation_Encoding_DefaultBinary), new(EUInformation))
	RegisterExtensionObject(NewNumericNodeID(0, id.ComplexNumberType_Encoding_DefaultBinary), new(ComplexNumberType))
	RegisterExtensionObject(NewNumericNodeID(0, id.DoubleComplexNumberType_Encoding_DefaultBinary), new(DoubleComplexNumberType))
	RegisterExtensionObject(NewNumericNodeID(0, id.AxisInformation_Encoding_DefaultBinary), new(AxisInformation))
	RegisterExtensionObject(NewNumericNodeID(0, id.XVType_Encoding_DefaultBinary), new(XVType))
	RegisterExtensionObject(NewNumericNodeID(0, id.ProgramDiagnosticDataType_Encoding_DefaultBinary), new(ProgramDiagnosticDataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.ProgramDiagnostic2DataType_Encoding_DefaultBinary), new(ProgramDiagnostic2DataType))
	RegisterExtensionObject(NewNumericNodeID(0, id.Annotation_Encoding_DefaultBinary), new(Annotation))
}
