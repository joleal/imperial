--CREATE USER TABLE
CREATE TABLE IF NOT EXISTS user (
    user_id varchar(255) NOT NULL primary key,
    email varchar(255) NOT NULL,
    nickname varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS game (
    game_id bigint unsigned default(uuid_short()) primary key,
    fk_game_settings bigint,
    fk_version int,
    name varchAR(255) NOT NULL,
    game_start bit NOT NULL DEFAULT 0,
    game_finish bit NOT NULL DEFAULT 0,
    current_country varchar(20) NULL,
    version varchar(20) NULL,
    number_of_players	INT NOT NULL,
    investor_card bit NOT NULL DEFAULT 1,
    tax_increase_only_bonus bit,
    starting_mode varchar(20)
);

CREATE TABLE IF NOT EXISTS game_version (
    id int primary key,
    name varchar(25)
);

CREATE TABLE IF NOT EXISTS game_player (
    id bigint unsigned default(uuid_short()) primary key,
    fk_game bigint,
    fk_user varchar(255),
	money    int,
	play_order    int,
	investor bool
);

CREATE TABLE IF NOT EXISTS game_settings (
    id bigint unsigned default(uuid_short()) primary key,
    fk_game bigint,
    fk_user bigint
);

ALTER TABLE game_player MODIFY fk_user VARCHAR(255);