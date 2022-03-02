Go Java Class File Parser
=========================

The jclass (package name `class`) parser support class files (those ending in  `.class`) as specified in [Chapter 4 of the Oracle JVM specification](http://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html). With the exception of the [Runtime[In]Visible[Paramterer]Annotations](http://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7.16), [AnnotationDefault](http://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7.20) & [StackMapTable](http://docs.oracle.com/javase/specs/jvms/se7/html/jvms-4.html#jvms-4.7.4) attributes. Otherwise all defined attributes & constants are supported and parsed correctly.

## Documentation

You can find the documentation [on GoDoc](http://godoc.org/github.com/jcla1/jclass). Additionally there are some [examples](examples/) provided in the repository.

## Use cases

First idea that comes to mind, is of course a JVM, but jclass can also be used for validating class files, obfuscating them or compressing them. This would be accomplished by, for example removing unnecessary LineNumberTable(s) and other attributes that don't affect the sematics of the class file, when executed.

## License

[MIT License](LICENSE)

### Changes after fork
    go get github.com/jcla1/jclass
    /Users/i321172/go/pkg/mod/github.com/jcla1/jclass@v0.0.0-20140603161227-b5cf858c0316/attributes.go:375:26: duplicate argument a
To fix above issue, folk the repo and do the changes. The repo is a bit old and last update in 2014...