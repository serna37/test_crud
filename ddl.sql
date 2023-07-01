-- public.mst_usr definition

-- Drop table

-- DROP TABLE mst_usr;

CREATE TABLE mst_usr (
	id serial4 NOT NULL, -- id
	"name" varchar(64) NOT NULL DEFAULT 'kitting'::character varying, -- 名前
	usr_login_id varchar(64) NOT NULL DEFAULT 'kitting'::character varying, -- ログインID
	usr_pass_word varchar(1024) NOT NULL DEFAULT 'kitting'::character varying, -- パスワード
	auth_token varchar(2048) NOT NULL DEFAULT 'kitting'::character varying, -- 認証トークン
	last_login timestamp NULL, -- 最終ログイン日時
	CONSTRAINT mst_usr_pkey PRIMARY KEY (id)
);
CREATE INDEX mst_usr_id_idx ON public.mst_usr USING btree (id);
COMMENT ON TABLE public.mst_usr IS 'ユーザ';

-- Column comments

COMMENT ON COLUMN public.mst_usr.id IS 'id';
COMMENT ON COLUMN public.mst_usr."name" IS '名前';
COMMENT ON COLUMN public.mst_usr.usr_login_id IS 'ログインID';
COMMENT ON COLUMN public.mst_usr.usr_pass_word IS 'パスワード';
COMMENT ON COLUMN public.mst_usr.auth_token IS '認証トークン';
COMMENT ON COLUMN public.mst_usr.last_login IS '最終ログイン日時';


-- public.trn_chat_room definition

-- Drop table

-- DROP TABLE trn_chat_room;

CREATE TABLE trn_chat_room (
	id bigserial NOT NULL, -- チャットルームID
	"name" varchar(64) NOT NULL DEFAULT 'chatroom'::character varying, -- ルーム名
	creater_id int4 NOT NULL DEFAULT 0, -- 作成者ID
	del_flg bool NOT NULL DEFAULT false, -- 削除フラグ
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 作成日時
	CONSTRAINT trn_chat_room_pk PRIMARY KEY (id)
);

-- Column comments

COMMENT ON COLUMN public.trn_chat_room.id IS 'チャットルームID';
COMMENT ON COLUMN public.trn_chat_room."name" IS 'ルーム名';
COMMENT ON COLUMN public.trn_chat_room.creater_id IS '作成者ID';
COMMENT ON COLUMN public.trn_chat_room.del_flg IS '削除フラグ';
COMMENT ON COLUMN public.trn_chat_room.created_at IS '作成日時';


-- public.mst_category definition

-- Drop table

-- DROP TABLE mst_category;

CREATE TABLE mst_category (
	id bigserial NOT NULL, -- カテゴリID
	usr_id int4 NOT NULL DEFAULT 0, -- ユーザID
	"name" varchar(64) NOT NULL DEFAULT 'no data'::character varying, -- カテゴリ名
	del_flg bool NOT NULL DEFAULT false, -- 削除フラグ
	CONSTRAINT mst_category_pk PRIMARY KEY (id),
	CONSTRAINT mst_category_fk FOREIGN KEY (usr_id) REFERENCES mst_usr(id)
);

-- Column comments

COMMENT ON COLUMN public.mst_category.id IS 'カテゴリID';
COMMENT ON COLUMN public.mst_category.usr_id IS 'ユーザID';
COMMENT ON COLUMN public.mst_category."name" IS 'カテゴリ名';
COMMENT ON COLUMN public.mst_category.del_flg IS '削除フラグ';


-- public.trn_contents definition

-- Drop table

-- DROP TABLE trn_contents;

CREATE TABLE trn_contents (
	id bigserial NOT NULL, -- コンテンツID
	usr_id int4 NOT NULL DEFAULT 0, -- ユーザID
	contents text NULL, -- コンテンツ
	category_id int4 NOT NULL DEFAULT 0, -- カテゴリID
	del_flg bool NOT NULL DEFAULT false, -- 削除フラグ
	title varchar(64) NULL, -- タイトル
	CONSTRAINT trn_contents_pk PRIMARY KEY (id),
	CONSTRAINT trn_contents_fk FOREIGN KEY (usr_id) REFERENCES mst_usr(id),
	CONSTRAINT trn_contents_fk_1 FOREIGN KEY (category_id) REFERENCES mst_category(id)
);

-- Column comments

COMMENT ON COLUMN public.trn_contents.id IS 'コンテンツID';
COMMENT ON COLUMN public.trn_contents.usr_id IS 'ユーザID';
COMMENT ON COLUMN public.trn_contents.contents IS 'コンテンツ';
COMMENT ON COLUMN public.trn_contents.category_id IS 'カテゴリID';
COMMENT ON COLUMN public.trn_contents.del_flg IS '削除フラグ';
COMMENT ON COLUMN public.trn_contents.title IS 'タイトル';


-- public.mst_tag definition

-- Drop table

-- DROP TABLE mst_tag;

CREATE TABLE mst_tag (
	id bigserial NOT NULL, -- タグID
	usr_id int4 NOT NULL DEFAULT 0, -- ユーザID
	"name" varchar(64) NOT NULL DEFAULT 'no data'::character varying, -- タグ名
	del_flg bool NOT NULL DEFAULT false, -- 削除フラグ
	category_id int4 NOT NULL DEFAULT 0, -- カテゴリID
	CONSTRAINT mst_tag_pk PRIMARY KEY (id),
	CONSTRAINT mst_tag_fk FOREIGN KEY (usr_id) REFERENCES mst_usr(id),
	CONSTRAINT mst_tag_fk_1 FOREIGN KEY (category_id) REFERENCES mst_category(id)
);

-- Column comments

COMMENT ON COLUMN public.mst_tag.id IS 'タグID';
COMMENT ON COLUMN public.mst_tag.usr_id IS 'ユーザID';
COMMENT ON COLUMN public.mst_tag."name" IS 'タグ名';
COMMENT ON COLUMN public.mst_tag.del_flg IS '削除フラグ';
COMMENT ON COLUMN public.mst_tag.category_id IS 'カテゴリID';


-- public.trn_contents_tag definition

-- Drop table

-- DROP TABLE trn_contents_tag;

CREATE TABLE trn_contents_tag (
	id bigserial NOT NULL, -- タグ付けID
	content_id int4 NOT NULL DEFAULT 0, -- コンテンツID
	tag_id int4 NOT NULL DEFAULT 0, -- タグID
	CONSTRAINT trn_contents_tag_pk PRIMARY KEY (id),
	CONSTRAINT trn_contents_tag_fk FOREIGN KEY (content_id) REFERENCES trn_contents(id),
	CONSTRAINT trn_contents_tag_fk_1 FOREIGN KEY (tag_id) REFERENCES mst_tag(id)
);

-- Column comments

COMMENT ON COLUMN public.trn_contents_tag.id IS 'タグ付けID';
COMMENT ON COLUMN public.trn_contents_tag.content_id IS 'コンテンツID';
COMMENT ON COLUMN public.trn_contents_tag.tag_id IS 'タグID';


-- public.trn_chat_room_member definition

-- Drop table

-- DROP TABLE trn_chat_room_member;

CREATE TABLE trn_chat_room_member (
	id bigserial NOT NULL, -- チャットルーム参加者管理ID
	chat_room_id int4 NOT NULL DEFAULT 0, -- チャットルームID
	joiner_id int4 NOT NULL DEFAULT 0, -- 参加者ID
	del_flg bool NOT NULL DEFAULT false, -- 削除フラグ
	joined_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 参加日時
	inviter int4 NOT NULL DEFAULT 0, -- 招待者
	CONSTRAINT trn_chat_room_member_pk PRIMARY KEY (id),
	CONSTRAINT trn_chat_room_member_fk FOREIGN KEY (chat_room_id) REFERENCES trn_chat_room(id)
);

-- Column comments

COMMENT ON COLUMN public.trn_chat_room_member.id IS 'チャットルーム参加者管理ID';
COMMENT ON COLUMN public.trn_chat_room_member.chat_room_id IS 'チャットルームID';
COMMENT ON COLUMN public.trn_chat_room_member.joiner_id IS '参加者ID';
COMMENT ON COLUMN public.trn_chat_room_member.del_flg IS '削除フラグ';
COMMENT ON COLUMN public.trn_chat_room_member.joined_at IS '参加日時';
COMMENT ON COLUMN public.trn_chat_room_member.inviter IS '招待者';


-- public.trn_chat_msg definition

-- Drop table

-- DROP TABLE trn_chat_msg;

CREATE TABLE trn_chat_msg (
	id bigserial NOT NULL, -- チャットルーム参加者管理ID
	chat_room_id int4 NOT NULL DEFAULT 0, -- チャットルームID
	from_id int4 NOT NULL DEFAULT 0, -- 発言者ID
	from_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 発言日時
	msg text NULL, -- メッセージ
	CONSTRAINT trn_chat_msg_pk PRIMARY KEY (id),
	CONSTRAINT trn_chat_msg_fk FOREIGN KEY (chat_room_id) REFERENCES trn_chat_room(id)
);

-- Column comments

COMMENT ON COLUMN public.trn_chat_msg.id IS 'チャットルーム参加者管理ID';
COMMENT ON COLUMN public.trn_chat_msg.chat_room_id IS 'チャットルームID';
COMMENT ON COLUMN public.trn_chat_msg.from_id IS '発言者ID';
COMMENT ON COLUMN public.trn_chat_msg.from_at IS '発言日時';
COMMENT ON COLUMN public.trn_chat_msg.msg IS 'メッセージ';

