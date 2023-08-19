# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/dynamic/v1/dynamic.proto](#proto_dynamic_v1_dynamic-proto)
    - [AddDynamicRequest](#proto-dynamic-v1-AddDynamicRequest)
    - [AddDynamicResponse](#proto-dynamic-v1-AddDynamicResponse)
    - [DeleteDynamicRequest](#proto-dynamic-v1-DeleteDynamicRequest)
    - [DeleteDynamicResponse](#proto-dynamic-v1-DeleteDynamicResponse)
    - [DynamicData](#proto-dynamic-v1-DynamicData)
    - [ListDynamicsRequest](#proto-dynamic-v1-ListDynamicsRequest)
    - [ListDynamicsResponse](#proto-dynamic-v1-ListDynamicsResponse)
    - [UpdateDynamicStatusRequest](#proto-dynamic-v1-UpdateDynamicStatusRequest)
    - [UpdateDynamicStatusResponse](#proto-dynamic-v1-UpdateDynamicStatusResponse)
  
    - [DynamicService](#proto-dynamic-v1-DynamicService)
  
- [proto/dynamic/v1/page.proto](#proto_dynamic_v1_page-proto)
    - [AddPageRequest](#proto-dynamic-v1-AddPageRequest)
    - [AddPageResponse](#proto-dynamic-v1-AddPageResponse)
    - [DeletePageRequest](#proto-dynamic-v1-DeletePageRequest)
    - [DeletePageResponse](#proto-dynamic-v1-DeletePageResponse)
    - [GetPagesRequest](#proto-dynamic-v1-GetPagesRequest)
    - [GetPagesResponse](#proto-dynamic-v1-GetPagesResponse)
    - [PageData](#proto-dynamic-v1-PageData)
    - [UpdatePageStatusRequest](#proto-dynamic-v1-UpdatePageStatusRequest)
    - [UpdatePageStatusResponse](#proto-dynamic-v1-UpdatePageStatusResponse)
  
    - [PageService](#proto-dynamic-v1-PageService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto_dynamic_v1_dynamic-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/dynamic/v1/dynamic.proto



<a name="proto-dynamic-v1-AddDynamicRequest"></a>

### AddDynamicRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| user_id | [string](#string) |  |  |






<a name="proto-dynamic-v1-AddDynamicResponse"></a>

### AddDynamicResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dynamic_id | [int32](#int32) |  |  |






<a name="proto-dynamic-v1-DeleteDynamicRequest"></a>

### DeleteDynamicRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dynamic_id | [int32](#int32) |  |  |






<a name="proto-dynamic-v1-DeleteDynamicResponse"></a>

### DeleteDynamicResponse







<a name="proto-dynamic-v1-DynamicData"></a>

### DynamicData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dynamic_id | [int32](#int32) |  |  |
| title | [string](#string) |  |  |
| overview | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| published | [bool](#bool) |  |  |






<a name="proto-dynamic-v1-ListDynamicsRequest"></a>

### ListDynamicsRequest







<a name="proto-dynamic-v1-ListDynamicsResponse"></a>

### ListDynamicsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dynamics | [DynamicData](#proto-dynamic-v1-DynamicData) | repeated |  |






<a name="proto-dynamic-v1-UpdateDynamicStatusRequest"></a>

### UpdateDynamicStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dynamic_id | [int32](#int32) |  |  |
| title | [string](#string) |  |  |
| overview | [string](#string) |  |  |
| published | [bool](#bool) |  |  |






<a name="proto-dynamic-v1-UpdateDynamicStatusResponse"></a>

### UpdateDynamicStatusResponse






 

 

 


<a name="proto-dynamic-v1-DynamicService"></a>

### DynamicService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListDynamics | [ListDynamicsRequest](#proto-dynamic-v1-ListDynamicsRequest) | [ListDynamicsResponse](#proto-dynamic-v1-ListDynamicsResponse) |  |
| AddDynamic | [AddDynamicRequest](#proto-dynamic-v1-AddDynamicRequest) | [AddDynamicResponse](#proto-dynamic-v1-AddDynamicResponse) |  |
| DeleteDynamic | [DeleteDynamicRequest](#proto-dynamic-v1-DeleteDynamicRequest) | [DeleteDynamicResponse](#proto-dynamic-v1-DeleteDynamicResponse) |  |
| UpdateDynamicStatus | [UpdateDynamicStatusRequest](#proto-dynamic-v1-UpdateDynamicStatusRequest) | [UpdateDynamicStatusResponse](#proto-dynamic-v1-UpdateDynamicStatusResponse) |  |

 



<a name="proto_dynamic_v1_page-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/dynamic/v1/page.proto



<a name="proto-dynamic-v1-AddPageRequest"></a>

### AddPageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| chapter_id | [int32](#int32) |  |  |
| order | [int32](#int32) |  |  |






<a name="proto-dynamic-v1-AddPageResponse"></a>

### AddPageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_id | [int32](#int32) |  |  |






<a name="proto-dynamic-v1-DeletePageRequest"></a>

### DeletePageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_id | [int32](#int32) |  |  |






<a name="proto-dynamic-v1-DeletePageResponse"></a>

### DeletePageResponse







<a name="proto-dynamic-v1-GetPagesRequest"></a>

### GetPagesRequest







<a name="proto-dynamic-v1-GetPagesResponse"></a>

### GetPagesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pages | [PageData](#proto-dynamic-v1-PageData) | repeated |  |






<a name="proto-dynamic-v1-PageData"></a>

### PageData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_id | [int32](#int32) |  |  |
| title | [string](#string) |  |  |
| order | [int32](#int32) |  |  |
| chapter_id | [int32](#int32) |  |  |






<a name="proto-dynamic-v1-UpdatePageStatusRequest"></a>

### UpdatePageStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_id | [int32](#int32) |  |  |
| title | [string](#string) |  |  |
| order | [int32](#int32) |  |  |






<a name="proto-dynamic-v1-UpdatePageStatusResponse"></a>

### UpdatePageStatusResponse






 

 

 


<a name="proto-dynamic-v1-PageService"></a>

### PageService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListPages | [GetPagesRequest](#proto-dynamic-v1-GetPagesRequest) | [GetPagesResponse](#proto-dynamic-v1-GetPagesResponse) |  |
| AddPage | [AddPageRequest](#proto-dynamic-v1-AddPageRequest) | [AddPageResponse](#proto-dynamic-v1-AddPageResponse) |  |
| DeletePage | [DeletePageRequest](#proto-dynamic-v1-DeletePageRequest) | [DeletePageResponse](#proto-dynamic-v1-DeletePageResponse) |  |
| UpdatePageStatus | [UpdatePageStatusRequest](#proto-dynamic-v1-UpdatePageStatusRequest) | [UpdatePageStatusResponse](#proto-dynamic-v1-UpdatePageStatusResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

