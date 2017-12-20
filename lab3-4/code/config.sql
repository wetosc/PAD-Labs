CREATE TABLE IF NOT EXISTS wine (
	id char(36) PRIMARY KEY,
	name TEXT,
	flavor TEXT,
	color TEXT,
	price FLOAT
);

CREATE TABLE IF NOT EXISTS cellar (
	id char(36) PRIMARY KEY,
	name TEXT,
	location TEXT,
	owner TEXT,
	area FLOAT
);


CREATE TABLE IF NOT EXISTS wine_to_cellar (
	wine_id char(36),
	cellar_id char(36),
	PRIMARY KEY (wine_id, cellar_id),
	FOREIGN KEY (wine_id) REFERENCES wine(id) ON DELETE CASCADE,
	FOREIGN KEY (cellar_id) REFERENCES cellar(id) ON DELETE CASCADE
);


