use proc_macro::TokenStream;
use quote::quote;
use syn::{parse_macro_input, AttributeArgs, Ident, ItemFn, Lit, NestedMeta};

#[proc_macro_attribute]
pub fn main(args: TokenStream, input: TokenStream) -> TokenStream {
    let input_path = match &parse_macro_input!(args as AttributeArgs)[..] {
        [NestedMeta::Lit(Lit::Int(day))] => format!("../../inputs/{}.in", day.token()),
        _ => panic!("Expected one integer argument"),
    };

    let mut aoc_solution = parse_macro_input!(input as ItemFn);
    aoc_solution.sig.ident = Ident::new("aoc_solution", aoc_solution.sig.ident.span());

    let tokens = quote! {
      const INPUT: &str = include_str!(#input_path);
      #aoc_solution
      fn main() {
        let now = ::std::time::Instant::now();
        let (p1, p2) = aoc_solution(INPUT.trim_end());
        let time = now.elapsed();
        println!("Part 1: {}", p1);
        println!("Part 2: {}", p2);
        println!("Time: {:.2?}", time);
      }
    };
    TokenStream::from(tokens)
}

#[proc_macro_attribute]
pub fn test(args: TokenStream, input: TokenStream) -> TokenStream {
    let day = match &parse_macro_input!(args as AttributeArgs)[..] {
        [NestedMeta::Lit(Lit::Int(day))] => day.token(),
        _ => panic!("Expected one integer argument"),
    };

    let input_path = format!("../../inputs/{}.test", day);
    let results_path = format!("../../inputs/{}.test.results", day);

    let mut aoc_test = parse_macro_input!(input as ItemFn);
    aoc_test.sig.ident = Ident::new("aoc_test", aoc_test.sig.ident.span());

    let tokens = quote! {
      const TEST_INPUT: &str = include_str!(#input_path);
      #aoc_test
      #[test]
      fn test() {
        let (p1, p2) = aoc_solution(TEST_INPUT.trim_end());
        let (r1, r2) = include_str!(#results_path).split_once("\r\n").unwrap();
        assert_eq!(p1, r1.parse::<usize>().unwrap());
        assert_eq!(p2, r2.parse::<usize>().unwrap());
      }
    };
    TokenStream::from(tokens)
}
