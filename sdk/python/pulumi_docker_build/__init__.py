# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .image import *
from .index import *
from .provider import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_docker_build.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_docker_build.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "docker-build",
  "mod": "index",
  "fqn": "pulumi_docker_build",
  "classes": {
   "docker-build:index:Image": "Image",
   "docker-build:index:Index": "Index"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "docker-build",
  "token": "pulumi:providers:docker-build",
  "fqn": "pulumi_docker_build",
  "class": "Provider"
 }
]
"""
)
