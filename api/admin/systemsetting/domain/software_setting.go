package domain

type SoftwareSettingRequest struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type SoftwareSetting struct {
	Id         int    `json:"id"`
	Content    string `json:"content"`
	Type       string `json:"type"`
	CreateUser int    `json:"createUser"`
	UpdateUser int    `json:"updateUser"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

func (SoftwareSetting) TableName() string {
	return "tb_software_setting"
}

/*
CREATE TABLE tal.tb_softeware_setting (
	id BIGINT UNSIGNED auto_increment NOT NULL,
	content TEXT NULL,
	create_time TIMESTAMP NULL,
	update_time TIMESTAMP NULL,
	create_user integer NOT NULL,
	update_user integer NULL,
	type varchar(20) NULL,
	CONSTRAINT tb_about_us_pk PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;

*/
