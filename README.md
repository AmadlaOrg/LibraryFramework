<img src=".assets/float.jpg" alt="Flying" style="width: 400px;" align="right">

# LibraryFramework 📚
📚 Framework | Library 📚

Using the decorator pattern it is possible to quickly add cli functionalities to an application.

Here is a simple example of how it can easily be implemented:
```go
func main() {
  cli.New(
    "hery",
    "HERY",
    "1.0.0",
    func(rootCmd *cobra.Command) {
      rootCmd.AddCommand(cmd.SettingsCmd)
      rootCmd.AddCommand(cmd.CollectionCmd)
    })
}
```

## ©️ Copyright
- "<a rel="noopener noreferrer" href="https://www.flickr.com/photos/37667416@N04/4031754170">1004115</a>" by <a rel="noopener noreferrer" href="https://www.flickr.com/photos/37667416@N04">Biblioteca Rector Machado y Nuñez</a> is marked with <a rel="noopener noreferrer" href="https://creativecommons.org/publicdomain/mark/1.0/?ref=openverse">Public Domain Mark 1.0 <img src="https://mirrors.creativecommons.org/presskit/icons/pd.svg" style="height: 1em; margin-right: 0.125em; display: inline;" /></a>.

## :scroll: License

The license for the code and documentation can be found in the [LICENSE](./LICENSE) file.

---

Made in Québec 🏴󠁣󠁡󠁱󠁣󠁿, Canada 🇨🇦!
