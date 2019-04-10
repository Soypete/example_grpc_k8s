CREATE TABLE houses
(
  taxvaluedollarcnt Integer,
  lotsizesquarefeet DECIMAL,
  yearbuilt INTEGER,
  bedroomcnt INTEGER,
  bathroomcnt DECIMAL,
  regionidcity TEXT,
  regionidzip TEXT,
);

COPY houses(
taxvaluedollarcnt,
yearbuilt,
bedroomcnt,
bathroomcnt,
lotsizesquarefeet,
regionidcity,
regionidzip)
FROM 'data/properties_2016.csv' DELIMITER ',' CSV HEADER;


