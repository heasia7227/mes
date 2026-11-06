"use client";

import { Row, Col } from "antd";
import DictionaryCatalogList from "./catalog";
import DictionaryDataList from "./data";

const DictionaryList = () => {
    return (
        <>
            <Row gutter={16}>
                <Col xs={24} md={8} lg={7} xl={6}>
                    <DictionaryCatalogList />
                </Col>
                <Col xs={24} md={16} lg={17} xl={18}>
                    <DictionaryDataList />
                </Col>
            </Row>
        </>
    );
};

export default DictionaryList;
