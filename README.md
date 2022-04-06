# Teora

Fan-made RPG based on [GeoEXE]'s [Gwain Saga].

NOTE: This project is neither affiliated with nor endorsed by GeoEXE.

## Building

Running `make` should build teora without any fuss.

If necessary, you can (re)build teora's assets with `make assets bento`.

However, it is recommended to build teora in debug mode using `make DEBUG=1` for now.
Among other things, it adds a TPS counter and builds teora as a console app so logs appear in the console (on Windows).

CGo is required, but cross-compiling from Linux to Windows works out of the box; The default Make target builds both a native and a Windows binary.

## Credits

GeoEXE for creating Gwain Sage in the first place.

The following external resources are vendorised:

- [Hack]: version 3.003

[GeoEXE]: https://www.youtube.com/c/geoexeofficial
[Gwain Saga]: https://youtube.com/playlist?list=PLtVNv5LHqiUMkdxa0eFlpZJEKxhyBzzr1
[Hack]: https://github.com/source-foundry/Hack
