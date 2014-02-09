# Copyright 2013 Google Inc. All Rights Reserved.

"""Command to print version information for Cloud SDK components.
"""

import textwrap

from google.cloud.sdk.calliope import base
from google.cloud.sdk.core import config
from google.cloud.sdk.core import log
from google.cloud.sdk.core.updater import update_manager


class Version(base.Command):
  """Print version information for Cloud SDK components."""

  def Run(self, args):
    manager = update_manager.UpdateManager()
    return manager.GetCurrentVersionsInformation()

  def Display(self, args, result):
    printables = []
    for name in sorted(result):
      version = result[name]
      printables.append('{name} {version}'.format(name=name, version=version))
    component_versions = '\n'.join(printables)
    log.Print(textwrap.dedent("""\
Google Cloud SDK {cloudsdk_version}
Copyright Google Inc. 2013.

{component_versions}
""".format(
    cloudsdk_version=config.CLOUD_SDK_VERSION,
    component_versions=component_versions)))
