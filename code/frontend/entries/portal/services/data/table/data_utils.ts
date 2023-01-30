const istclass = ["shorttext", "phonenumber", "number", "select"];

export const generate_column_order = (columns: { [_: string]: object }): string[] => {
  const doneCols = {};
  const orderedColumns = [];

  // then first class colums
  istclass.forEach((cType) => {
    Object.values(columns).forEach((val) => {
      if (doneCols[val["slug"]]) {
        return;
      }

      if (val["ref_type"]) {
        return;
      }

      if (val["ctype"] !== cType) {
        return;
      }

      orderedColumns.push(val["slug"]);
      doneCols[val["slug"]] = true;
    });
  });

  // then remaining columns expect ref types
  Object.values(columns).forEach((val) => {
    if (istclass.includes(val["ctype"])) {
      return;
    }
    if (doneCols[val["slug"]]) {
      return;
    }

    if (val["ref_type"]) {
      return;
    }

    orderedColumns.push(val["slug"]);
    doneCols[val["slug"]] = true;
  });

  // atlast ref types
  Object.values(columns).forEach((val) => {
    if (doneCols[val["slug"]]) {
      return;
    }

    if (!val["ref_type"]) {
      return;
    }

    orderedColumns.push(val["slug"]);
    doneCols[val["slug"]] = true;
  });

  return orderedColumns;
};