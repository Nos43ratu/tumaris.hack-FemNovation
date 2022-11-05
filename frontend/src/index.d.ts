declare global {
  type Product = {
    id: number;
    name: string;
    description: string;
    product_categories: string;
    product_sizes: string;
    colors: string;
    weight: string;
    product_comments: string;
    rating: number;
    price: string;
    imageSrc: string;
  };
}

export {};
