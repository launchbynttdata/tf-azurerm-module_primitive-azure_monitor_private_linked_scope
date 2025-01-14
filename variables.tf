// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

# COMMON
variable "name" {
  description = "Name of the Private Link Scope"
  type        = string
}

variable "resource_group_name" {
  description = "Resource group name"
  type        = string
}

variable "ingestion_access_mode" {
  description = "The ingestion access mode for the Azure Monitor Private Link Scope."
  type        = string
  default     = null
}

variable "query_access_mode" {
  description = "The query access mode for the Azure Monitor Private Link Scope."
  type        = string
  default     = null
}

variable "tags" {
  description = "Custom tags for the Private Link Scope"
  type        = map(string)
  default     = {}
}

# Private Link Scope inputs

variable "linked_resource_ids" {
  description = "Map of resources to associate with the Private Link Scope"
  type        = map(string)
  default     = {}
}
