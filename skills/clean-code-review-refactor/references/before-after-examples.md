# Clean code before/after examples

Use these as pattern examples for reviews and refactors.

## 1. Tight coupling inside, loose coupling outside

### Before: page owns every technical detail

A React page imports storage, transport, analytics, and rendering details directly:

- component calls `fetch()` inline
- component writes to `localStorage`
- component assembles analytics payloads
- component transforms raw server data before rendering

### After: page composes a smaller domain-facing surface

Split into layers:

- page composes the use case and passes domain-friendly props
- `useOrdersPage()` or similar hook coordinates data flow
- API adapter handles transport details
- storage helper handles persistence details
- analytics helper handles event emission

Result: internal pieces may still be tightly coordinated, but the page surface stays loose and readable.

## 2. Hide abstraction in the implementation

### Before: callers must understand plumbing

A UI component exposes props like:

- `transportAdapter`
- `cacheProvider`
- `queryFactory`
- `repositoryImpl`

The caller must understand plumbing to use the component.

### After: callers speak domain intent

Expose intent-oriented props instead:

- `orderId`
- `onApprove`
- `mode`

Move plumbing behind a hook, service, or composition boundary.

Result: the user-facing API speaks the domain, not the implementation.

## 3. Single Responsibility Principle

### Before: the app shell does everything

`App.tsx`:

- defines routes
- sets up providers
- loads boot data
- holds feature flags
- contains screen-specific conditional rendering
- stores large config objects inline

### After: bootstrap, routes, providers, and config are split

Split into:

- `app/providers.tsx`
- `app/routes.tsx`
- `app/bootstrap.ts`
- `app/config/*.ts`
- feature-level screen modules

Result: each file explains itself instead of becoming a central accident archive.

## 4. Config outside logic-heavy files

### Before: static maps are buried in component logic

A component contains:

- inline button variant maps
- feature flag mappings
- status-to-label dictionaries
- route metadata objects

mixed with rendering logic and event handlers.

### After: config modules carry the static knowledge

Move static configuration into dedicated files such as:

- `feature/config/status-labels.ts`
- `feature/config/button-variants.ts`
- `feature/config/route-meta.ts`

Result: render code becomes easier to scan and config evolves independently.

## 5. React layering instead of one giant tree

### Before: every layer is mixed into one component tree

A single page component mixes:

- app shell layout
- section layout
- UI primitives
- widget composition
- async state branching
- domain actions
- modal orchestration

### After: route, state, widgets, and primitives are separated

Separate by layer:

- route or screen file for high-level composition
- hook for state and actions
- feature widgets for domain sections
- presentational components for reusable UI blocks
- primitives stay dumb and low-level

Result: less nesting, clearer ownership, easier reuse, and smaller review surfaces.
