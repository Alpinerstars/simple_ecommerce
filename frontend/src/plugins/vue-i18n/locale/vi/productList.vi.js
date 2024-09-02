const productList = {
  label: {
    title: 'List of all products',
    resultCount: 'matching results',
  },
  filter: {
    sort: {
      title: 'Sort',
      placeholder: 'Default',
      options: [
        {
          value: 'product_updated_at.desc',
          label: 'Latest',
        },
        {
          value: 'product_sold.desc',
          label: 'Best Sellers',
        },
        {
          value: 'product_price.asc',
          label: 'Price Ascending',
        },
        {
          value: 'product_price.desc',
          label: 'Price Descending',
        },
      ],
      optionsLength: '4',
    },
    brand: {
      title: 'Brand',
      placeholder: 'All',
    },
  },
};
export default productList;
