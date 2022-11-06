declare global {
  type Product = {
    id: number;
    shop_id: number;
    name: string;
    description: string;
    sizes: string[];
    colors: number[];
    weight: number;
    price: number;
    category_id: number;
    rating: 9.8;
  };

  type ApiResponse<Data> = {
    data: {
      error: null | string;
      response: {
        data: Data;
        status: number;
      };
    };
  };

  type UserData = {
    id: number;
    email: string;
    phone_number: string;
    first_name: string;
    last_name: string;
    password: string;
    role: string;
    about_me: string;
    instagram: string;
    rating: number;
    shop_id: number;
  };

  type Order = {
    id: number;
    client_id: number;
    shop_id: number;
    status: OrderStatus | number;
    product_id: number;
    products: Product;
    created_at: {
      Time: string;
      Valid: boolean;
    };
    payed_at: {
      Time: string;
      Valid: boolean;
    };
    packed_at: {
      Time: string;
      Valid: boolean;
    };
    delivered_at: {
      Time: string;
      Valid: boolean;
    };
    cancel_reason: string;
  };

  enum OrderStatus {
    "0",
    "1",
    "2",
    "3",
    "4",
    "5",
    "-10",
    "-1",
    "-2",
  }
}

export {};
