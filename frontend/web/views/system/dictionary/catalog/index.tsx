"use client";
import { Card, Button, Tree, Popconfirm } from "antd";
import DictionaryCatalogAdd from "./add";
import DictionaryCatalogDelete from "./delete";

const DictionaryCatalogList = () => {
    return (
        <>
            <Card size="small" className="main-card">
                <div
                    style={{
                        display: "flex",
                        justifyContent: "space-between",
                        marginBottom: 16,
                    }}
                >
                    <div style={{ fontSize: 16, fontWeight: 500 }}>字典目录</div>
                    <div>
                        <DictionaryCatalogAdd />
                    </div>
                </div>
                <Tree
                    blockNode
                    defaultExpandAll
                    showLine
                    treeData={[]}
                    // selectedKeys={selectedCatalogId ? [selectedCatalogId] : []}
                    // onSelect={(keys) => setSelectedCatalogId((keys[0] as string) ?? undefined)}
                    style={{ height: "calc(100vh - 206px)" }}
                />

                <div style={{ marginTop: 16, display: "flex", gap: 8 }}>
                    <DictionaryCatalogAdd isEdit={true} />
                    <DictionaryCatalogDelete />
                </div>
            </Card>
        </>
    );
};

export default DictionaryCatalogList;
