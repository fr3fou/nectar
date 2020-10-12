# local-booru

A place for weebs to save, organize and share their saved pics.

- Save pics from anywhere (drag & drop)
- Link to source (mandatory?)
- Organize collections by tags (fetch tags from other boorus, ability to add custom)
- Share collections with unique URL
- Link / Follow collections (automatically add pics from your collection and vice versa)
    - Support mutual collections 
- Neat overview of random pics
- Search by tag

- Collection
    - Has many images
    - Like a folder

- Image
    - Has many tags
    - Belongs to many collections
    - Link to S3 CDN

- User
    - Has many collections
