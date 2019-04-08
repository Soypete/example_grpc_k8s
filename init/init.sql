CREATE TABLE houses
(
  id serial NOT NULL,
  price Integer,
  lot_size DECIMAL,
  age INTEGER,
  number_of_bedrooms INTEGER,
  number_of_bathrooms DECIMAL,
  city TEXT,
  state TEXT,
  zip TEXT,
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


