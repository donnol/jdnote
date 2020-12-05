-- 其实只需要指定用户名和表名，还有该表的特有字段就可以生成以下内容

-- 新建t_file_content表
CREATE TABLE public.t_file_content (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL,
    updated_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL,
	created_by integer NOT NULL DEFAULT 0::integer,
	updated_by integer NOT NULL DEFAULT 0::integer,

    -- 特有字段
    content bytea NOT NULL DEFAULT ''::bytea
);

-- 更新时间绑定触发器
CREATE TRIGGER update_updated_at BEFORE INSERT OR UPDATE ON public.t_file_content
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

ALTER TABLE public.t_file_content OWNER TO jd;

-- 创建序列
CREATE SEQUENCE public.t_file_content_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE public.t_file_content_id_seq OWNER TO jd;

ALTER SEQUENCE public.t_file_content_id_seq OWNED BY public.t_file_content.id;

-- 设置自增
ALTER TABLE ONLY public.t_file_content ALTER COLUMN id SET DEFAULT nextval('public.t_file_content_id_seq'::regclass);

-- 设置主键
ALTER TABLE ONLY public.t_file_content
    ADD CONSTRAINT t_file_content_pkey PRIMARY KEY (id);
