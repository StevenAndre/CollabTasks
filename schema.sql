
CREATE TABLE users(

    id_user varchar(4) not null,
    name varchar(60) not null,
    lastname varchar(70) not null,
    username varchar(60)not null unique,
    email varchar(200) not null unique,
    password varchar(255)not null,
    profile_photo varchar(300),
    CONSTRAINT pk_users PRIMARY KEY(id_user)

);

CREATE TABLE notifications(
    id_notification varchar(10)not null,
    issue varchar(150) not null,
    message varchar(300)not null,
    issued TIMESTAMP not null,
    id_user varchar(4) not null,
    CONSTRAINT pk_notifications PRIMARY KEY(id_notification),
    CONSTRAINT fk_user_notofocations FOREIGN KEY (id_user) REFERENCES users(id_user) ON DELETE CASCADE
    
);

CREATE TABLE projects(
    id_project varchar(8)not null,
    name varchar(100) not null,
    description varchar(250),
    created_at TIMESTAMP NOT NULL,
    create_by varchar(4) not null,
    CONSTRAINT pk_projects PRIMARY KEY(id_project),
    CONSTRAINT fk_user_projects FOREIGN KEY (create_by) REFERENCES users(id_user) ON DELETE CASCADE
    
);

CREATE TABLE tasks(
    id_task varchar(10)not null,
    title varchar(100) not null,
    description varchar(250) not null,
    created_at TIMESTAMP not null,
    iscompleted boolean default false,
    id_user varchar(4) not null,
    id_project varchar(8)not null,
    CONSTRAINT pk_tasks PRIMARY KEY(id_task),
    CONSTRAINT fk_user_tasks FOREIGN KEY (id_user) REFERENCES users(id_user),
    CONSTRAINT fk_project_tasks FOREIGN KEY (id_project) REFERENCES projects(id_project)

);


CREATE TABLE evidences(
    id_evidence varchar(10)not null,
    filepath varchar(300) not null,
    comment varchar(200),
    task_id varchar(10) not null,
    created_at TIMESTAMP not null,
    CONSTRAINT pk_evidences PRIMARY KEY(id_evidence),
    CONSTRAINT fk_task_evidence FOREIGN KEY (task_id) REFERENCES tasks(task_id)

);

CREATE TABLE positions(
    id_position SMALLINT not null,
    name VARCHAR(50) not null,
    description VARCHAR(100),
    CONSTRAINT pk_positions PRIMARY KEY(id_position)
);

CREATE TABLE teams(
    id_team varchar(8) not null,
    name VARCHAR(100) not null,
    CONSTRAINT pk_teams PRIMARY KEY(id_team),
);

CREATE TABLE team_members(
    id_member varchar(4) not null,
    id_team varchar(8) not null,
    position_id SMALLINT not null,
    CONSTRAINT pk_team_members PRIMARY KEY(id_member,id_team),
    CONSTRAINT fk_team_members FOREIGN KEY (id_team) REFERENCES teams(id_team),
    CONSTRAINT fk_users_team_members FOREIGN KEY (id_member) REFERENCES users(id_user),
    CONSTRAINT fk_position_team_members FOREIGN KEY (position_id) REFERENCES positions(id_position),
);

CREATE TABLE assigned_project(
    id_project varchar(8) not null,
    id_team varchar(8) not null,
    CONSTRAINT pk_assigned_projects PRIMARY KEY(id_project,id_team),
    CONSTRAINT fk_project_assignments FOREIGN KEY (id_project) REFERENCES projects(id_project),
    CONSTRAINT fk_team_assignments FOREIGN KEY (id_team) REFERENCES teams(id_team)

);