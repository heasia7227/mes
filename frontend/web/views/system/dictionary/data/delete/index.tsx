"use client";

import { Button } from "antd";
import { SysDictionary } from "@/lib/types/system/dictionary";

interface IProps {
    record?: SysDictionary;
    refresh?: () => void;
}

const DictionaryDataDelete = ({ record }: IProps) => {
    return (
        <>
            <Button type="link" size="small" danger>
                删除
            </Button>
        </>
    );
};

export default DictionaryDataDelete;
