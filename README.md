# Teora

Fan-made RPG based on [GeoEXE]'s [Gwain Saga].

NOTE: This project is neither affiliated with nor endorsed by GeoEXE.

## Building

Running `make` should build teora, along with any assets, without any fuss.

CGo is required, but cross-compiling from Linux to Windows works out of the box; The default Make target builds both a native and a Windows binary.

For debugging, you can `make debug` instead. Among other things, it builds teora as a console app so logs appear in the console (on Windows).

## Credits

GeoEXE for creating Gwain Sage in the first place.

The following external resources are vendorised:

- [Hack]: version 3.003

[GeoEXE]: https://www.youtube.com/c/geoexeofficial
[Gwain Saga]: https://youtube.com/playlist?list=PLtVNv5LHqiUMkdxa0eFlpZJEKxhyBzzr1
[Hack]: https://github.com/source-foundry/Hack
