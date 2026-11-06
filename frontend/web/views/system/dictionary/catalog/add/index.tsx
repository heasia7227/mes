"use client";

import { Button } from "antd";
import { FolderAddOutlined, EditOutlined } from "@ant-design/icons";
import { SysDictionaryCatalog } from "@/lib/types/system/dictionary-catalog";

interface IProps {
    isEdit?: boolean;
    record?: SysDictionaryCatalog;
    refresh?: () => void;
}

const DictionaryCatalogAdd = ({ isEdit = false, record, refresh }: IProps) => {
    return (
        <>
            {isEdit ? (
                <Button size="small" icon={<EditOutlined />} disabled={!record} block>
                    编辑
                </Button>
            ) : (
                <Button type="link" size="small" icon={<FolderAddOutlined />}>
                    新增
                </Button>
            )}
        </>
    );
};

export default DictionaryCatalogAdd;
