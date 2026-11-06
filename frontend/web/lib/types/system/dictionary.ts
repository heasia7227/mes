import { BaseEntity } from "@/lib/types/types";

/** t_sys_dictionary */
export interface SysDictionary extends BaseEntity {
    catalog_id: string;
    catalog_code: string;
    dictionary_value: string;
    dictionary_name: string;
    sort?: number;
    association_dict?: string;
    association_dict_name?: string;
    association_dict_value?: string;
}
