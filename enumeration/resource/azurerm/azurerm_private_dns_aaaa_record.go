package azurerm

import (
	"github.com/snyk/driftctl/enumeration/resource"
)

const AzurePrivateDNSAAAARecordResourceType = "azurerm_private_dns_aaaa_record"

func initAzurePrivateDNSAAAARecordMetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetHumanReadableAttributesFunc(AzurePrivateDNSAAAARecordResourceType, func(res *resource.Resource) map[string]string {
		val := res.Attrs
		attrs := make(map[string]string)
		if name := val.GetString("name"); name != nil && *name != "" {
			attrs["Name"] = *name
		}
		if zone := val.GetString("zone_name"); zone != nil && *zone != "" {
			attrs["Zone"] = *zone
		}
		return attrs
	})
	resourceSchemaRepository.SetFlags(AzurePrivateDNSAAAARecordResourceType, resource.FlagDeepMode)
}
