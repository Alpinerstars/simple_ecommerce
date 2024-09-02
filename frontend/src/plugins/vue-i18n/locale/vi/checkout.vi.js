const checkout = {
  cartList: {
    header: {
      product: 'Product',
      price: 'Unit price',
      quantity: 'Quantity',
      total: 'Total price',
    },
  },
  cartItem: {
    multiple: 'x',
    incQuantity: '+',
    decQuantity: '-',
  },
  cartTotals: {
    title: 'Order Total',
    content: {
      totals: 'Total Amount',
    },
    additionalInfo: 'The above prices include tax and shipping.',
    submit: 'Place Order',
  },
  deliveryAddress: {
    title: 'Delivery Address',
    content: {
      noAddress: {
        title:
          'There is currently no information about the delivery address, please add another address.',
        addNewAddress: 'Add an address',
      },
    },
  },
};

export default checkout;
