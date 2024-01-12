CREATE TABLE "data_keluhan"(
    "id_keluhan" SERIAL,
    "kategori_keluhan" VARCHAR(255) NOT NULL,
    "id_user" BIGINT NOT NULL,
    "keluhan" TEXT NOT NULL,
    "id_dokter_handler" BIGINT NOT NULL,
    "status" VARCHAR(255),
        "createdAt" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
        "updatedAt" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "data_keluhan" ADD PRIMARY KEY("id_keluhan");
CREATE TABLE "data_dokter"(
    "id_dokter" SERIAL,
    "name" VARCHAR(255) NOT NULL,
    "occupation" VARCHAR(255) NOT NULL,
    "workexperience" VARCHAR(255) NOT NULL,
    "phonenumber" BIGINT NOT NULL,
    "createdAt" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updatedAt" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "data_dokter" ADD PRIMARY KEY("id_dokter");
ALTER TABLE
    "data_keluhan" ADD CONSTRAINT "data_keluhan_id_dokter_handler_foreign" FOREIGN KEY("id_dokter_handler") REFERENCES "data_dokter"("id_dokter");
INSERT INTO data_dokter VALUES(2,'Dr Novan','Pro','asadada','9819283932','2011-01-01 00:00:00','2011-01-01 00:00:00')