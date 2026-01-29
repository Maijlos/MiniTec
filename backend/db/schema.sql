CREATE TABLE Project (
  id   BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name text    NOT NULL,
  code text    NOT NULL
);

CREATE TABLE Station (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name text NOT NULL,
  project_id BIGINT NOT NULL,
  FOREIGN KEY (project_id) REFERENCES Project(id) ON DELETE CASCADE
);

CREATE TABLE State (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    final_state INT NOT NULL,
    start_date DATETIME,
    end_date DATETIME,
    station_id BIGINT NOT NULL,
    FOREIGN KEY (station_id) REFERENCES Station(id) ON DELETE CASCADE
);

CREATE TABLE User (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    first_name text NOT NULL,
    second_name text NOT NULL,
    email text NOT NULL,
    isAdmin BOOLEAN NOT NULL DEFAULT FALSE,
    isBanned BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE User_Project (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    project_id BIGINT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES User(id),
    FOREIGN KEY (project_id) REFERENCES Project(id) ON DELETE CASCADE
);