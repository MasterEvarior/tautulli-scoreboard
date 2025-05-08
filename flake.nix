{
  description = "Development flake for Tautulli Scoreboard";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.11";
  };

  outputs =
    { nixpkgs, ... }:
    let
      x86 = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages."${x86}";
    in
    {
      devShells."${x86}".default = pkgs.mkShellNoCC {
        packages = with pkgs; [
          # Golang
          go
          golangci-lint

          # Formatters
          treefmt
          mdformat
          yamlfmt
          jsonfmt
          deadnix
          nixfmt-rfc-style
        ];

        shellHook = ''
          git config --local core.hooksPath .githooks/
        '';

        # Environment Variables
        TS_BASE_URL = "";
        TS_API_TOKEN = "";
      };
    };
}
