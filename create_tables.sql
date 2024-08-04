DROP TABLE IF EXISTS statistics;
CREATE TABLE statistics (
  id        INT AUTO_INCREMENT NOT NULL,
  age_min	INT	NOT NULL,
  age_max	INT,
  age      	VARCHAR(128) NOT NULL,
  salary    VARCHAR(255) NOT NULL,
  price     DECIMAL(7,2) NOT NULL,
  race 		VARCHAR(128) NOT NULL,
  height 	INT NOT NULL,
  is_married BOOLEAN,
  ip 		VARCHAR(15),
  date_created DATE NOT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS unique_ips_temp;
CREATE TABLE unique_ips_temp (
  id        INT AUTO_INCREMENT NOT NULL,
  ip 		VARCHAR(15) NOT NULL,
  date_created DATE NOT NULL,
  PRIMARY KEY (`id`)
);
