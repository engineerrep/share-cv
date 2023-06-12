SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- 数据库： `app_share_cv`
--

CREATE DATABASE IF NOT EXISTS app_share_cv DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;
use app_share_cv;

-- app语言表
DROP TABLE IF EXISTS `app_language`;
CREATE TABLE `app_language`
(
    `id`                bigint       NOT NULL COMMENT '编号' AUTO_INCREMENT,
    `country_code`      varchar(10) NOT NULL DEFAULT '' COMMENT '国家图标',
    `full_code`         varchar(20)  NOT NULL DEFAULT '' COMMENT '代码全称',
    `code`              varchar(10)  NOT NULL DEFAULT '' COMMENT '代码简称',
    `name`              varchar(100) NOT NULL DEFAULT '' COMMENT '语言',
    `name_en`           varchar(100) NOT NULL DEFAULT '' COMMENT '语言英文',
    `name_cn`           varchar(100) NOT NULL DEFAULT '' COMMENT '语言中文',
    `currency`          varchar(10)  NOT NULL DEFAULT '' COMMENT '币种',
    `create_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at`         timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`        timestamp             DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY(`id`),
    UNIQUE KEY `full_code` (`full_code`),
    KEY        `idx_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='app语言表';


INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('en-au','English(Australia)','English(Australia)','英语(澳大利亚)','AU','en','AUD',1);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('en-ca','English(Canada)','English(Canada)','英语(加拿大)','CA','en','CAD',2);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('en-in','English(India)','English(India)','英语(印度)','IN','en','INR',3);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('en-ie','English(Ireland)','English(Ireland)','英语(爱尔兰)','IE','en','EUR',4);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('en-my','English(Malaysia)','English(Malaysia)','英语(马来西亚)','MY','en','MYR',5);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('en-nz','English(New Zealand)','English(New Zealand)','英语(新西兰)','NZ','en','NZD',6);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('en-sg','English(Singapore)','English(Singapore)','英语(新加坡)','SG','en','SGD',7);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('en-za','English(South Africa)','English(South Africa)','英语(南非)','ZA','en','ZAR',8);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('en-gb','English(United Kingdom)','English(United Kingdom)','英语(英国)','GB','en','GBP',9);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('en-us','English(United States)','English(United States)','英语(美国)','US','en','USD',10);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('es-mx','Español(México)','Spanish(Mexico)','西班牙语(墨西哥)','MX','es','MXN',11);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('id-id','Bahasa Indonesia','Indonesian','印度尼西亚语','ID','id','IDR',12);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('ms-my','Bahasa Melayu','Malay(Malaysia)','马来语(马来西亚)','MY','ms','MYR',13);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('bs-latn-ba','Bosanski','Bosnian(Latin)','波斯尼亚语(拉丁语系)','BA','bs','BAM',14);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('ca-es','Català','Catalan','加泰罗尼亚语','ES','ca','EUR',15);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('cs-cz','Čeština','Czech','捷克语','CZ','cs','CZK',16);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('da-dk','Dansk','Danish','丹麦语','DK','da','DKK',17);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('de-at','Deutsch(Österreich)','German(Austria)','德语(奥地利)','AT','de','EUR',18);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('de-ch','Deutsch(Schweiz)','German(Switzerland)','德语(瑞士)','CH','de','CHF',19);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('de-de','Deutsch','German','德语','DE','de','EUR',20);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('et-ee','Eesti','Estonian','爱沙尼亚语','EE','et','EUR',21);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('es-es','Español','Spanish','西班牙语','ES','es','EUR',22);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('eu-es','Euskara','Basque','巴斯克语','ES','eu','EUR',23);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('fil-ph','Filipino','Filipino','菲律宾语','PH','fil','PHP',24);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('fr-be','Français(Belgique)','French(Belgium)','法语(比利时)','BE','fr','EUR',25);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('fr-ca','Français(Canada)','French(Canada)','法语(加拿大)','CA','fr','CAD',26);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('fr-ch','Français(Suisse)','French(Switzerland)','法语(瑞士)','CH','fr','CHF',27);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('fr-fr','Français','French','法语','FR','fr','EUR',28);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('ga-ie','Gaeilge','Irish','爱尔兰语','IE','ga','EUR',29);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('gl-es','Galego','Galician','加利西亚语','ES','gl','EUR',30);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('hr-hr','Hrvatski','Croatian','克罗地亚语','HR','hr','HRK',31);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('is-is','Íslenska','Icelandic','冰岛语','IS','is','ISK',32);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('it-ch','Italiano(Svizzera)','Italian(Switzerland)','意大利语(瑞士)','CH','it','CHF',33);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('it-it','Italiano','Italian','意大利语','IT','it','EUR',34);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('lv-lv','Latviešu','Latvian','拉脱维亚语','LV','lv','EUR',35);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('lb-lu','Lëtzebuergesch','Luxembourgish','卢森堡语','LU','lb','EUR',36);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('lt-lt','Lietuvių','Lithuanian','立陶宛语','LT','lt','EUR',37);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('hu-hu','Magyar','Hungarian','匈牙利语','HU','hu','HUF',38);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('mt-mt','Malti','Maltese','马耳他语','MT','mt','EUR',39);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('nl-be','Nederlands(België)','Dutch(Belgium)','荷兰语(比利时)','BE','nl','EUR',40);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('nl-nl','Nederlands','Dutch','荷兰语','NL','nl','EUR',41);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('nb-no','Norsk Bokmål','Norwegian Bokmål','书面挪威语','NO','nb','NOK',42);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('pl-pl','Polski','Polish','波兰语','PL','pl','PLN',43);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('pt-br','Português(Brasil)','Portuguese(Brazil)','葡萄牙语(巴西)','BR','pt','BRL',44);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('pt-pt','Português(Portugal)','Portuguese(Portugal)','葡萄牙语(葡萄牙)','PT','pt','EUR',45);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('ro-ro','Română','Romanian','罗马尼亚语','RO','ro','RON',46);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('sk-sk','Slovenčina','Slovak','斯洛伐克语','SK','sk','EUR',47);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('sl-si','Slovenski','Slovenian','斯洛文尼亚语','SI','sl','EUR',48);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('sr-latn-rs','Srpski','Serbian(Latin)','塞尔维亚语(拉丁)','RS','sr','RSD',49);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('fi-fi','Suomi','Finnish','芬兰语','FI','fi','EUR',50);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('sv-se','Svenska','Swedish','瑞典语','SE','sv','SEK',51);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('vi-vn','TiếngViệt','Vietnamese','越南语','VN','vi','VND',52);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('tr-tr','Türkçe','Turkish','土耳其语','TR','tr','TRY',53);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('el-gr','Ελληνικά','Greek','希腊语','GR','el','EUR',54);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('bs-cyrl-ba','Босански','Bosnian(Cyrillic)','波斯尼亚语(西里尔文)','BA','bs','BAM',55);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('bg-bg','Български','Bulgarian','保加利亚语','BG','bg','BGN',56);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('kk-kz','қазақ тілі','Kazakh','哈萨克语','KZ','kk','KZT',57);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('ru-ru','Русский','Russian','俄语','RU','ru','RUB',58);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('sr-cyrl-rs','Српски','Serbian(Cyrillic)','塞尔维亚语(西里尔)','RS','sr','RSD',59);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('uk-ua','Українська','Ukrainian','乌克兰语','UA','uk','UAH',60);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('he-il','עברית','Hebrew','希伯来语','IL','he','ILS',61);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('ar-sa','العربية','Arabic','阿拉伯语','SA','ar','SAR',62);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('hi-in','हिंदी','Hindi','印地语','IN','hi','INR',63);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('th-th','ไทย','Thai','泰语','TH','th','THB',64);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('ko-kr','한국어','Korean','韩语','KR','ko','KRW',65);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('ja-jp','日本語','Japanese','日语','JP','ja','JPY',66);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('zh-cn','中文(简体)','Chinese(Simplified)','中文(简体)','CN','zh','CNY',67);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('zh-tw','中文(繁体 台湾)','Chinese(Traditional)','中文(繁体 台湾)','TW','zh','TWD',68);
INSERT INTO `app_language`(`full_code`, `name`, `name_en`, `name_cn`,`country_code`,`code`,`currency`,`id`) VALUE('zh-hk','中文(繁体 香港)','Chinese(Traditional, Hong Kong)','中文(繁体 香港)','HK','zh','HKD',69);

