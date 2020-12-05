
-- 新建{{.TableName}}表
CREATE TABLE public.{{.TableName}} (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL,
    updated_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL,
    created_by integer NOT NULL DEFAULT 0::integer,
    updated_by integer NOT NULL DEFAULT 0::integer,

    -- 特有字段
    {{.SpecialFields}}
);

-- 更新时间绑定触发器
CREATE TRIGGER update_updated_at BEFORE INSERT OR UPDATE ON public.{{.TableName}}
    FOR EACH ROW EXECUTE FUNCTION update_updated_at();

ALTER TABLE public.{{.TableName}} OWNER TO {{.UserName}};

-- 创建序列
CREATE SEQUENCE public.{{.TableName}}_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE public.{{.TableName}}_id_seq OWNER TO {{.UserName}};

ALTER SEQUENCE public.{{.TableName}}_id_seq OWNED BY public.{{.TableName}}.id;

-- 设置自增
ALTER TABLE ONLY public.{{.TableName}} ALTER COLUMN id SET DEFAULT nextval('public.{{.TableName}}_id_seq'::regclass);

-- 设置主键
ALTER TABLE ONLY public.{{.TableName}}
    ADD CONSTRAINT {{.TableName}}_pkey PRIMARY KEY (id);
