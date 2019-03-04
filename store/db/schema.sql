--
-- PostgreSQL database dump
--

-- Dumped from database version 10.7 (Ubuntu 10.7-1.pgdg16.04+1)
-- Dumped by pg_dump version 10.6 (Ubuntu 10.6-1.pgdg16.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: t_action; Type: TABLE; Schema: public; Owner: jd
--

CREATE TABLE public.t_action (
    id bigint NOT NULL,
    action text NOT NULL
);


ALTER TABLE public.t_action OWNER TO jd;

--
-- Name: t_action_id_seq; Type: SEQUENCE; Schema: public; Owner: jd
--

CREATE SEQUENCE public.t_action_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_action_id_seq OWNER TO jd;

--
-- Name: t_action_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jd
--

ALTER SEQUENCE public.t_action_id_seq OWNED BY public.t_action.id;


--
-- Name: t_note; Type: TABLE; Schema: public; Owner: jd
--

CREATE TABLE public.t_note (
    id bigint NOT NULL,
    user_id integer NOT NULL,
    detail text NOT NULL,
    created_at timestamp with time zone DEFAULT clock_timestamp() NOT NULL,
    title text DEFAULT ''::text NOT NULL
);


ALTER TABLE public.t_note OWNER TO jd;

--
-- Name: t_note_id_seq; Type: SEQUENCE; Schema: public; Owner: jd
--

CREATE SEQUENCE public.t_note_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_note_id_seq OWNER TO jd;

--
-- Name: t_note_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jd
--

ALTER SEQUENCE public.t_note_id_seq OWNED BY public.t_note.id;


--
-- Name: t_role; Type: TABLE; Schema: public; Owner: jd
--

CREATE TABLE public.t_role (
    id bigint NOT NULL,
    role text NOT NULL
);


ALTER TABLE public.t_role OWNER TO jd;

--
-- Name: t_role_action; Type: TABLE; Schema: public; Owner: jd
--

CREATE TABLE public.t_role_action (
    id bigint NOT NULL,
    role_id bigint NOT NULL,
    action_id bigint NOT NULL
);


ALTER TABLE public.t_role_action OWNER TO jd;

--
-- Name: t_role_action_id_seq; Type: SEQUENCE; Schema: public; Owner: jd
--

CREATE SEQUENCE public.t_role_action_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_role_action_id_seq OWNER TO jd;

--
-- Name: t_role_action_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jd
--

ALTER SEQUENCE public.t_role_action_id_seq OWNED BY public.t_role_action.id;


--
-- Name: t_role_id_seq; Type: SEQUENCE; Schema: public; Owner: jd
--

CREATE SEQUENCE public.t_role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_role_id_seq OWNER TO jd;

--
-- Name: t_role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jd
--

ALTER SEQUENCE public.t_role_id_seq OWNED BY public.t_role.id;


--
-- Name: t_user; Type: TABLE; Schema: public; Owner: jd
--

CREATE TABLE public.t_user (
    id bigint NOT NULL,
    phone character(11) NOT NULL,
    password character varying(256) NOT NULL,
    name character varying(256) NOT NULL,
    email text DEFAULT ''::text NOT NULL
);


ALTER TABLE public.t_user OWNER TO jd;

--
-- Name: t_user_id_seq; Type: SEQUENCE; Schema: public; Owner: jd
--

CREATE SEQUENCE public.t_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_user_id_seq OWNER TO jd;

--
-- Name: t_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jd
--

ALTER SEQUENCE public.t_user_id_seq OWNED BY public.t_user.id;


--
-- Name: t_user_role; Type: TABLE; Schema: public; Owner: jd
--

CREATE TABLE public.t_user_role (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    role_id bigint NOT NULL
);


ALTER TABLE public.t_user_role OWNER TO jd;

--
-- Name: t_user_role_id_seq; Type: SEQUENCE; Schema: public; Owner: jd
--

CREATE SEQUENCE public.t_user_role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.t_user_role_id_seq OWNER TO jd;

--
-- Name: t_user_role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: jd
--

ALTER SEQUENCE public.t_user_role_id_seq OWNED BY public.t_user_role.id;


--
-- Name: t_action id; Type: DEFAULT; Schema: public; Owner: jd
--

ALTER TABLE ONLY public.t_action ALTER COLUMN id SET DEFAULT nextval('public.t_action_id_seq'::regclass);


--
-- Name: t_note id; Type: DEFAULT; Schema: public; Owner: jd
--

ALTER TABLE ONLY public.t_note ALTER COLUMN id SET DEFAULT nextval('public.t_note_id_seq'::regclass);


--
-- Name: t_role id; Type: DEFAULT; Schema: public; Owner: jd
--

ALTER TABLE ONLY public.t_role ALTER COLUMN id SET DEFAULT nextval('public.t_role_id_seq'::regclass);


--
-- Name: t_role_action id; Type: DEFAULT; Schema: public; Owner: jd
--

ALTER TABLE ONLY public.t_role_action ALTER COLUMN id SET DEFAULT nextval('public.t_role_action_id_seq'::regclass);


--
-- Name: t_user id; Type: DEFAULT; Schema: public; Owner: jd
--

ALTER TABLE ONLY public.t_user ALTER COLUMN id SET DEFAULT nextval('public.t_user_id_seq'::regclass);


--
-- Name: t_user_role id; Type: DEFAULT; Schema: public; Owner: jd
--

ALTER TABLE ONLY public.t_user_role ALTER COLUMN id SET DEFAULT nextval('public.t_user_role_id_seq'::regclass);


--
-- Name: t_action t_action_pkey; Type: CONSTRAINT; Schema: public; Owner: jd
--

ALTER TABLE ONLY public.t_action
    ADD CONSTRAINT t_action_pkey PRIMARY KEY (id);


--
-- Name: t_note t_note_pkey; Type: CONSTRAINT; Schema: public; Owner: jd
--

ALTER TABLE ONLY public.t_note
    ADD CONSTRAINT t_note_pkey PRIMARY KEY (id);


--
-- Name: t_role_action t_role_action_pkey; Type: CONSTRAINT; Schema: public; Owner: jd
--

ALTER TABLE ONLY public.t_role_action
    ADD CONSTRAINT t_role_action_pkey PRIMARY KEY (id);


--
-- Name: t_role t_role_pkey; Type: CONSTRAINT; Schema: public; Owner: jd
--

ALTER TABLE ONLY public.t_role
    ADD CONSTRAINT t_role_pkey PRIMARY KEY (id);


--
-- Name: t_user t_user_pkey; Type: CONSTRAINT; Schema: public; Owner: jd
--

ALTER TABLE ONLY public.t_user
    ADD CONSTRAINT t_user_pkey PRIMARY KEY (id);


--
-- Name: t_user_role t_user_role_pkey; Type: CONSTRAINT; Schema: public; Owner: jd
--

ALTER TABLE ONLY public.t_user_role
    ADD CONSTRAINT t_user_role_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

