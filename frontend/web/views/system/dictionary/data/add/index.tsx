"use client";

import { Button } from "antd";
import { PlusOutlined, EditOutlined } from "@ant-design/icons";
import { SysDictionary } from "@/lib/types/system/dictionary";

interface IProps {
    isEdit?: boolean;
    record?: SysDictionary;
    refresh?: () => void;
}

const DictionaryDataAdd = ({ isEdit = false, record, refresh }: IProps) => {
    return (
        <>
            {isEdit ? (
                <Button type="link" size="small" icon={<EditOutlined />}>
                    编辑
                </Button>
            ) : (
                <Button type="primary" size="large" icon={<PlusOutlined />}>
                    新增
                </Button>
            )}
        </>
    );
};

export default DictionaryDataAdd;
