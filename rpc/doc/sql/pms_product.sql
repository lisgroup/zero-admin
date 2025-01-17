create table pms_product
(
    id                            bigint auto_increment
        primary key,
    brand_id                      bigint         not null,
    product_category_id           bigint         not null,
    feight_template_id            bigint         not null,
    product_attribute_category_id bigint         not null,
    name                          varchar(64)    not null,
    pic                           varchar(255)   not null,
    product_sn                    varchar(64)    not null comment '货号',
    delete_status                 int(1)         not null comment '删除状态：0->未删除；1->已删除',
    publish_status                int(1)         not null comment '上架状态：0->下架；1->上架',
    new_status                    int(1)         not null comment '新品状态:0->不是新品；1->新品',
    recommand_status              int(1)         not null comment '推荐状态；0->不推荐；1->推荐',
    verify_status                 int(1)         not null comment '审核状态：0->未审核；1->审核通过',
    sort                          int            not null comment '排序',
    sale                          int            not null comment '销量',
    price                         decimal(10, 2) not null,
    promotion_price               decimal(10, 2) not null comment '促销价格',
    gift_growth                   int default 0  not null comment '赠送的成长值',
    gift_point                    int default 0  not null comment '赠送的积分',
    use_point_limit               int            not null comment '限制使用的积分数',
    sub_title                     varchar(255)   not null comment '副标题',
    description                   text           not null comment '商品描述',
    original_price                decimal(10, 2) not null comment '市场价',
    stock                         int            not null comment '库存',
    low_stock                     int            not null comment '库存预警值',
    unit                          varchar(16)    not null comment '单位',
    weight                        decimal(10, 2) not null comment '商品重量，默认为克',
    preview_status                int(1)         not null comment '是否为预告商品：0->不是；1->是',
    service_ids                   varchar(64)    not null comment '以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮',
    keywords                      varchar(255)   not null,
    note                          varchar(255)   not null,
    album_pics                    varchar(255)   not null comment '画册图片，连产品图片限制为5张，以逗号分割',
    detail_title                  varchar(255)   not null,
    detail_desc                   text           not null,
    detail_html                   text           not null comment '产品详情网页内容',
    detail_mobile_html            text           not null comment '移动端网页详情',
    promotion_start_time          datetime       not null comment '促销开始时间',
    promotion_end_time            datetime       not null comment '促销结束时间',
    promotion_per_limit           int            not null comment '活动限购数量',
    promotion_type                int(1)         not null comment '促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购',
    brand_name                    varchar(255)   not null comment '品牌名称',
    product_category_name         varchar(255)   not null comment '商品分类名称',
    product_category_id_Array     varchar(64)   not null comment  '商品分类id字符串'
)
    comment '商品信息';

INSERT INTO pms_product (id, brand_id, product_category_id, feight_template_id, product_attribute_category_id, name, pic, product_sn, delete_status, publish_status, new_status, recommand_status, verify_status, sort, sale, price, promotion_price, gift_growth, gift_point, use_point_limit, sub_title, description, original_price, stock, low_stock, unit, weight, preview_status, service_ids, keywords, note, album_pics, detail_title, detail_desc, detail_html, detail_mobile_html, promotion_start_time, promotion_end_time, promotion_per_limit, promotion_type, brand_name, product_category_name) VALUES (33, 6, 35, 0, 1, '小米（MI）小米电视4A ', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5b02804dN66004d73.jpg', '4609652', 0, 1, 0, 0, 0, 0, 0, 2499.00, 100.00, 0, 0, 0, '小米（MI）小米电视4A 55英寸 L55M5-AZ/L55M5-AD 2GB+8GB HDR 4K超高清 人工智能网络液晶平板电视', '', 2499.00, 100, 0, '', 0.00, 0, '', '', '', '', '', '', '', '', '2021-03-16 20:50:48', '2021-03-16 20:50:56', 0, 0, '小米', '手机数码','1,2');
INSERT INTO pms_product (id, brand_id, product_category_id, feight_template_id, product_attribute_category_id, name, pic, product_sn, delete_status, publish_status, new_status, recommand_status, verify_status, sort, sale, price, promotion_price, gift_growth, gift_point, use_point_limit, sub_title, description, original_price, stock, low_stock, unit, weight, preview_status, service_ids, keywords, note, album_pics, detail_title, detail_desc, detail_html, detail_mobile_html, promotion_start_time, promotion_end_time, promotion_per_limit, promotion_type, brand_name, product_category_name) VALUES (34, 6, 35, 0, 1, '小米（MI）小米电视4A 65英寸', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5b028530N51eee7d4.jpg', '4609660', 0, 1, 0, 0, 0, 0, 0, 3999.00, 100.00, 0, 0, 0, ' L65M5-AZ/L65M5-AD 2GB+8GB HDR 4K超高清 人工智能网络液晶平板电视', '', 3999.00, 100, 0, '', 0.00, 0, '', '', '', '', '', '', '', '', '2021-03-16 20:50:48', '2021-03-16 20:50:57', 0, 0, '小米', '手机数码','1,2');
INSERT INTO pms_product (id, brand_id, product_category_id, feight_template_id, product_attribute_category_id, name, pic, product_sn, delete_status, publish_status, new_status, recommand_status, verify_status, sort, sale, price, promotion_price, gift_growth, gift_point, use_point_limit, sub_title, description, original_price, stock, low_stock, unit, weight, preview_status, service_ids, keywords, note, album_pics, detail_title, detail_desc, detail_html, detail_mobile_html, promotion_start_time, promotion_end_time, promotion_per_limit, promotion_type, brand_name, product_category_name) VALUES (35, 58, 29, 0, 11, '耐克NIKE 男子 休闲鞋 ROSHE RUN 运动鞋 511881-010黑色41码', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5b235bb9Nf606460b.jpg', '6799342', 0, 1, 0, 0, 0, 0, 0, 369.00, 100.00, 0, 0, 0, '耐克NIKE 男子 休闲鞋 ROSHE RUN 运动鞋 511881-010黑色41码', '', 369.00, 100, 0, '', 0.00, 0, '', '', '', '', '', '', '', '', '2021-03-16 20:50:48', '2021-03-16 20:50:58', 0, 0, 'NIKE', '男鞋','1,2');
INSERT INTO pms_product (id, brand_id, product_category_id, feight_template_id, product_attribute_category_id, name, pic, product_sn, delete_status, publish_status, new_status, recommand_status, verify_status, sort, sale, price, promotion_price, gift_growth, gift_point, use_point_limit, sub_title, description, original_price, stock, low_stock, unit, weight, preview_status, service_ids, keywords, note, album_pics, detail_title, detail_desc, detail_html, detail_mobile_html, promotion_start_time, promotion_end_time, promotion_per_limit, promotion_type, brand_name, product_category_name) VALUES (36, 58, 29, 0, 11, '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', 'http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/5b19403eN9f0b3cb8.jpg', '6799345', 0, 1, 1, 1, 0, 0, 0, 499.00, 100.00, 0, 0, 0, '耐克NIKE 男子 气垫 休闲鞋 AIR MAX 90 ESSENTIAL 运动鞋 AJ1285-101白色41码', '', 499.00, 100, 0, '', 0.00, 0, '', '', '', '', '', '', '', '', '2021-03-16 20:50:48', '2021-03-16 20:50:59', 0, 0, 'NIKE', '男鞋','1,2');
