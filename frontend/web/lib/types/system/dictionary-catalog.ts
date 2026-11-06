import { BaseEntity } from "@/lib/types/types";

/** t_sys_dictionary_catalog */
export interface SysDictionaryCatalog extends BaseEntity {
    catalog_name: string;
    catalog_code: string;
    parent_id?: string;
}
