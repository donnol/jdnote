-- 新建t_file表
CREATE TABLE public.t_file (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL,
    updated_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL,
	created_by integer NOT NULL DEFAULT 0::integer,
	updated_by integer NOT NULL DEFAULT 0::integer,

    name text DEFAULT ''::text NOT NULL,
    size integer NOT NULL,
	file_content_id integer NOT NULL
);

-- 更新时间绑定触发器
CREATE TRIGGER update_updated_at BEFORE INSERT OR UPDATE ON public.t_file
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

ALTER TABLE public.t_file OWNER TO jd;

-- 创建序列
CREATE SEQUENCE public.t_file_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE public.t_file_id_seq OWNER TO jd;

ALTER SEQUENCE public.t_file_id_seq OWNED BY public.t_file.id;

-- 设置自增
ALTER TABLE ONLY public.t_file ALTER COLUMN id SET DEFAULT nextval('public.t_file_id_seq'::regclass);

-- 设置主键
ALTER TABLE ONLY public.t_file
    ADD CONSTRAINT t_file_pkey PRIMARY KEY (id);
