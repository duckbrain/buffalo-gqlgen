# Buffalo GQLGen Plugn

## Plans

- Buffalo plugin
  - Generate the integration
    - Handler from package
    - Config from built-in template
    - Resolver file with built-in scalars pre-defined
  - Rerun gqlgen generate on rebuild (if needed)
  - generate resource includes schema and resolvers
- GQLGen plugin
  - Add custom directives
  - Auto generate pop resolvers (@pop directive?)
  - Auto bind common/Buffalo gql types (UUID, Time, nulls.X)
  - Bind db transaction on each mutation call
- Buffalo http handler for gqlgen
- Marshaller/Unmarshaller for common/buffalo scalar types

Directives can be added for permissions
