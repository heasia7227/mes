"use client";

import { Button } from "antd";
import { DeleteOutlined } from "@ant-design/icons";
import { SysDictionaryCatalog } from "@/lib/types/system/dictionary-catalog";

interface IProps {
    record?: SysDictionaryCatalog;
    refresh?: () => void;
}

const DictionaryCatalogDelete = ({ record, refresh }: IProps) => {
    return (
        <>
            <Button size="small" danger icon={<DeleteOutlined />} disabled={!record} block>
                删除
            </Button>
        </>
    );
};

export default DictionaryCatalogDelete;
