# Buffalo GQLGen Plugn

## Plans

- Buffalo plugin
  - Generate the integration
    - Handler from package
  - Rerun gqlgen generate on rebuild (if needed)
- GQLGen plugin
  - Add custom directives
  - Auto generate pop resolvers (@pop directive?)
  - Auto bind common/Buffalo gql types (UUID, Time, nulls.X)
  - Bind db transaction on each mutation call
