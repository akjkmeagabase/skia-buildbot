#!/usr/bin/env python
# Copyright (c) 2013 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

""" Run the Skia GM executable. """

from build_step import BuildStep
import build_step
import os
import sys


JSON_SUMMARY_FILENAME = 'actual-results.json'


class RunGM(BuildStep):
  def _Run(self):
    output_dir = os.path.join(self._device_dirs.GMActualDir(),
                              self._builder_name)
    cmd = ['--verbose',
           '--writeChecksumBasedFilenames',
           # Don't bother writing out image files that match our expectations--
           # we know that previous runs have already uploaded those!
           '--mismatchPath', output_dir,
           '--missingExpectationsPath', output_dir,
           '--writeJsonSummaryPath', os.path.join(output_dir,
                                                  JSON_SUMMARY_FILENAME),
           '--ignoreErrorTypes',
               'IntentionallySkipped', 'MissingExpectations',
               'ExpectationsMismatch',
           '--resourcePath', self._device_dirs.ResourceDir(),
           ] + self._gm_args

    device_gm_expectations_path = self._flavor_utils.DevicePathJoin(
        self._device_dirs.GMExpectedDir(), build_step.GM_EXPECTATIONS_FILENAME)
    if self._flavor_utils.DevicePathExists(device_gm_expectations_path):
      cmd.extend(['--readPath', device_gm_expectations_path])

    device_ignore_failures_path = self._flavor_utils.DevicePathJoin(
        self._device_dirs.GMExpectedDir(),
        build_step.GM_IGNORE_FAILURES_FILE)
    if self._flavor_utils.DevicePathExists(device_ignore_failures_path):
      cmd.extend(['--ignoreFailuresFile', device_ignore_failures_path])

    if 'Xoom' in self._builder_name:
      # The Xoom's GPU will crash on some tests if we don't use this flag.
      # http://code.google.com/p/skia/issues/detail?id=1434
      cmd.append('--resetGpuContext')

    # Exercise alternative renderModes, but not on the slowest platforms.
    # See https://code.google.com/p/skia/issues/detail?id=1641 ('Run GM tests
    # with all rendering modes enabled, SOMETIMES')
    if (not 'Android' in self._builder_name and
        not 'ChromeOS' in self._builder_name):
      cmd.extend(['--deferred', '--pipe', '--replay', '--rtree', '--serialize',
                  '--tileGrid'])

    if sys.platform == 'darwin':
      # msaa16 is flaky on Macs (driver bug?) so we skip the test for now
      cmd.extend(['--config', 'defaults', '~msaa16'])
    elif ('RazrI' in self._builder_name or
          'Nexus10' in self._builder_name or
          'Nexus4' in self._builder_name):
      cmd.extend(['--config', 'defaults', 'msaa4'])
    elif (not 'NoGPU' in self._builder_name and
          not 'ChromeOS' in self._builder_name and
          not 'GalaxyNexus' in self._builder_name):
      cmd.extend(['--config', 'defaults', 'msaa16'])

    if 'ZeroGPUCache' in self._builder_name:
      cmd.extend(['--gpuCacheSize', '0', '0', '--config', 'gpu'])
    if 'Nexus4' in self._builder_name:
      cmd.extend(['--useDocumentInsteadOfDevice',
                  '--forcePerspectiveMatrix',
                  # Disabling the following tests because they crash GM in
                  # perspective mode.
                  # See https://code.google.com/p/skia/issues/detail?id=1665
                  '--match',
                  '~scaled_tilemodes',
                  '~convexpaths',
                  '~clipped-bitmap'])

    self._flavor_utils.RunFlavoredCmd('gm', cmd)


if '__main__' == __name__:
  sys.exit(BuildStep.RunBuildStep(RunGM))
