CREATE OR REPLACE FUNCTION to_date_and_compare1(text, date)
RETURNS boolean AS $$
BEGIN
  RETURN to_date($1, 'YYYY-MM-DD') < $2;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION to_date_and_compare2(text, date)
RETURNS boolean AS $$
BEGIN
  RETURN to_date($1, 'YYYY-MM-DD') > $2;
END;
$$ LANGUAGE plpgsql;

CREATE OPERATOR |< (
  LEFTARG = text,
  RIGHTARG = date,
  PROCEDURE = to_date_and_compare1,
  COMMUTATOR = |>
);

CREATE OPERATOR |> (
  LEFTARG = text,
  RIGHTARG = date,
  PROCEDURE = to_date_and_compare2,
  COMMUTATOR = |<
);